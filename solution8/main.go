package main

import (
	"fmt"
)

func main() {
	var (
		n int64
		i int
	)
	fmt.Scan(&n, &i)
	i--
	if n&(1<<i) == 1<<i {
		fmt.Println(n &^ (1 << i)) // Change to 0 if on i-th index was 1
	} else {
		fmt.Println(n | (1 << i)) // Change to 1 if on i-th index was 0
	}
}
