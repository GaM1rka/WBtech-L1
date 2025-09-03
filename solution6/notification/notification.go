package main

import (
	"fmt"
	"time"
)

func worker(stopChan chan string) {
	for {
		select {
		case msg := <-stopChan:
			if msg != "" {
				fmt.Println("Get a notification to stop goroutine.")
				return
			}
		default:
			fmt.Println("Goroutine is working..")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stopChan := make(chan string)
	go worker(stopChan)

	time.Sleep(3 * time.Second)
	stopChan <- "stop"
	time.Sleep(1 * time.Second)
}
