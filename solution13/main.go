package main

import (
	"fmt"
)

func change(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	a := 5
	b := 3
	fmt.Printf("Before changing a = %d, b = %d\n", a, b)
	change(&a, &b)
	fmt.Printf("After changing a = %d, b = %d\n", a, b)
}
