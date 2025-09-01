package main

import (
	"bufio"
	"os"
	"strconv"
)

func pow(number int, channel chan int) {
	channel <- number * number // Storing results into channel
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	writer := bufio.NewWriter(os.Stdout)
	channel := make(chan int) // Channel to store output data

	for _, el := range array {
		go pow(el, channel) // Running function pow in goroutine
	}

	for range 5 {
		res := <-channel // Getting results from channel
		writer.WriteString(strconv.Itoa(res) + "\n")
	}

	// Reset buffer
	writer.Flush()
}
