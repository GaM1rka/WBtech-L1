package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 49, 52, 76, 8, 3, 10}
	var i int
	fmt.Print("Enter the index that you want to delete: ")
	fmt.Scan(&i)

	copy(slice[i:], slice[i+1:]) // Avoid creation new basic arrays.
	slice = slice[:len(slice)-1] // Deleting last element.

	fmt.Println(slice)
}
