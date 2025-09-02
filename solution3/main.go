package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, channel chan string) {
	defer wg.Done()
	for message := range channel {
		fmt.Printf("Worker %d read message: %s\n", id, message) // Reading message from channel
	}
}

func main() {
	var (
		n       int
		wg      sync.WaitGroup
		channel chan string = make(chan string)
	)
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, &wg, channel)
	}

	ctx, cancel := context.WithCancel(context.Background()) // Context variable for managing the goroutines work.

	go func() {
		i := 0
		for {
			i++
			select {
			case <-ctx.Done():
				fmt.Println("Writter stopped")
			case channel <- fmt.Sprintf("message number: %d", i):
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)

	fmt.Println("Stopping the writer after 3 seconds of working...")
	cancel() // Give a sign what would writer goroutine stop to work.

	close(channel) // Close channel to stop the workers
	wg.Wait()      // Waite while all worker finish to read from channel
}
