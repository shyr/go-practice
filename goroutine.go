package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "http://www.example.com"
	queue <- "http://www.example.net"
	queue <- "http://www.example.com/foo"
	queue <- "http://www.example.com/bar"

}

func fetchURL(queue chan string) {
	for {
		url, more := <-queue
		if more {
			// url 취득 처리
			fmt.Println("fetching", url)
			// ...
		} else {
			fmt.Println("worker exit")
			wg.Done()
			return
		}

	}
}
