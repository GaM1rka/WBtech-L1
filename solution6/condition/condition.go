package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) {
	cnt := 0
	for {
		fmt.Println("Goroutine is working..")
		time.Sleep(1 * time.Second)
		cnt++
		if cnt >= 5 {
			fmt.Println("Goroutine stop to work.")
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(&wg)
	wg.Wait()
}
