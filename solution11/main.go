package main

import "fmt"

func intersection(a, b []int) []int {
	first := make(map[int]bool)

	for _, el := range a {
		first[el] = true
	}

	res := make([]int, 0)
	for _, el := range b {
		if _, ok := first[el]; ok == true {
			res = append(res, el)
		}
	}

	return res
}

func main() {
	fmt.Println(intersection([]int{1, 2, 3}, []int{2, 3, 4}))
}
