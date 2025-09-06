package main

import (
	"fmt"
	"sync"
)

var Second int = 10e6 // Let assume that Golang execute 10e6 operations per second

// Function that is make sleeping of the whole program using simple for loop
func sleep(cnt int) {
	temp := 0
	for i := 0; i < cnt*Second; i++ {
		temp++
	}
}

func worker(sleepChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var flag bool
	for {
		select {
		case <-sleepChan:
			fmt.Printf("Worker sleeps for %d seconds\n", 5)
			sleep(5) // Make a sleeping
			flag = true
		default:
			fmt.Println("Worker is working...")
			if flag == true {
				return
			}
		}
	}

}

func main() {
	var wg sync.WaitGroup
	var sleepChan chan string = make(chan string)

	wg.Add(1)
	go worker(sleepChan, &wg)
	sleep(1) // Give a time to start the goroutine
	sleepChan <- "sleep"

	wg.Wait()
}
