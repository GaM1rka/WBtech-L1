package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker() {
	defer fmt.Println("This will be printed, because it's called with defer")

	runtime.Goexit()

	fmt.Println("This will not be printed, because written after runtime.Goexit()")
}

func main() {
	go worker()
	time.Sleep(1 * time.Second)
}
