package main

import (
	"fmt"
)

func readArray(array []int, arrayCh chan int) {
	for i := 0; i < len(array); i++ {
		arrayCh <- array[i]
	}
	close(arrayCh) // Closing after writing data to channel
}
func doubleArray(arrayCh chan int, doubleCh chan int) {
	for msg := range arrayCh {
		doubleCh <- msg * 2
	}
	close(doubleCh) // Closing after writing data to channel
}

func main() {
	var (
		arrayCh  chan int = make(chan int)
		doubleCh chan int = make(chan int)
	)

	array := []int{52, 4, 2, 3, 1, 7}

	go readArray(array, arrayCh)
	go doubleArray(arrayCh, doubleCh)

	for msg := range doubleCh {
		fmt.Println("Message from doubleCh: ", msg)
	}
}
