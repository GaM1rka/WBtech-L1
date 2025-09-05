package main

import (
	"fmt"
)

func createHugeString(size int) string {
	s := ""
	for i := 0; i < size; i++ {
		s += "a"
	}
	return s
}

func someFunc() string {
	v := createHugeString(1 << 10)
	res := string([]byte(v[:100])) // Create new string, to clear used storage of large string
	return res
}

func main() {
	justString := someFunc() // making local variabe
	fmt.Println(justString)
}
