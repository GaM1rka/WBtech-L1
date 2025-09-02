package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	defer cancel()

	go func() {
		i := 0
		for {
			i++
			select {
			case <-ctx.Done():
				fmt.Println("Writter stopped")
				return
			case channel <- fmt.Sprintf("message number: %d", i):
				continue
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Program is running, press Ctrl+C to stop")
	<-sigChan

	cancel()
	close(channel) // Close channel to stop the workers
	wg.Wait()      // Waite while all worker finish to read from channel
}
