package main

import (
	"context"
	"fmt"
	"time"
)

func reader(channel chan string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Reader stopped.")
			return
		case msg, ok := <-channel:
			if !ok {
				return
			}
			fmt.Printf("Readed message: %s", msg)
		}
	}
}

func writer(channel chan string, ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Writer stopped.")
			return
		case channel <- fmt.Sprintf("message number: %d.\n", i):
			i++
		}
	}
}

func main() {
	var (
		n       int
		channel chan string = make(chan string)
	)
	fmt.Scan(&n)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go writer(channel, ctx)
	go reader(channel, ctx)

	time.Sleep(time.Duration(n) * time.Second)

	cancel()
	time.Sleep(100 * time.Millisecond)
	close(channel)
}
