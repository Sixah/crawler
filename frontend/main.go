package main

import (
	"crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/home/shuai/ws/go/src/crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("/home/shuai/ws/go/src/crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
