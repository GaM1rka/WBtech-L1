package main

import (
	"fmt"
)

func unique(words []string) []string {
	res := make([]string, 0)
	wordMap := make(map[string]bool)

	for _, el := range words {
		wordMap[el] = true
	}

	for key, _ := range wordMap {
		res = append(res, key)
	}

	return res
}

func main() {
	fmt.Println(unique([]string{"cat", "cat", "dog", "cat", "tree"}))
}
