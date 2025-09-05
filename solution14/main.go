package main

import (
	"fmt"
)

func define(variable interface{}) {
	switch v := variable.(type) {
	case int:
		v *= 2
		fmt.Printf("Variable type is int and it's doubled: %d\n", v)
	case string:
		v += ", world!"
		fmt.Printf("Variable type is string and it's changed: %s\n", v)
	case bool:
		v = !v
		fmt.Println("Variable type is bool and it's changed: %v\n", v)
	case chan int:
		fmt.Printf("Variable type is chan int and message from chan is: %d\n", <-v)
	default:
		fmt.Println("Type of this variable is unknown!")
	}
}

func main() {
	var temp interface{} = 5
	define(temp)
}
