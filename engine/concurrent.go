package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	// ConfigureMasterWorkerChan(chan Request)
	// WorkerReady(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 保存重复的url数据
	// file, err := os.Create("url.html")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	//
	// write := bufio.NewWriter(file)
	// defer write.Flush()

	// in := make(chan Request)
	out := make(chan ParseResult)
	// e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		// createWorker(in, out)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			// log.Printf("Got item #%d: %v", itemCount, item)
			// itemCount++

			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				log.Printf("url repeat: %v", request.Url)
				// fmt.Fprintf(write, "%s\n", request.Url) // 保存重复的url数据
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

// func createWorker(in chan Request, out chan ParseResult) {
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	// in := make(chan Request)
	go func() {
		for {
			// tell scheduler I'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
