package main

import (
	"fmt"
	"gogogo/advanced/interface/retriever/mock"
	real2 "gogogo/advanced/interface/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "This is a fake imooc.com"}
	fmt.Printf("%T %v\n", r, r) // mock.Retriever {This is a fake imooc.com}

	r = real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	fmt.Printf("%T %v\n", r, r) // real.Retriever {Mozilla/5.0 1m0s}

	//fmt.Println(download(r))
}
