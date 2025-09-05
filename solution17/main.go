package main

import (
	"fmt"
)

func bsearch(arr []int, target int) int {
	l, r := 0, len(arr)-1

	for l+1 < r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 10, 15, 49, 52, 77}
	target := 10000
	fmt.Println(bsearch(arr, target))
}
