package main

import (
	"fmt"
	"sync"
)

type counter struct {
	cnt int
	mu  sync.Mutex
}

func increment(c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt++
}

func main() {
	var c counter
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&c, &wg)
	}

	wg.Wait()
	fmt.Println(c.cnt)
}
