package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("sixah")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun/shanghai/",
	// 	ParserFunc: parser.ParserCity,
	// })
}
