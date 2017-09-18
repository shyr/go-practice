package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ { // 2개의 goroutine을 생성
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "http://www.example.com"
	queue <- "http://www.example.net"
	queue <- "http://www.example.com/foo"
	queue <- "http://www.example.com/bar"

	close(queue) // goroutine에게 종료를 전달
	wg.Wait()    // 모든 goroutine이 종료되는 것을 대기
}

func fetchURL(queue chan string) {
	for {
		url, more := <-queue
		// fmt.Println("more: " + more)
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
