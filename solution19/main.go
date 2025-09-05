package main

import "fmt"

func reverse(s []rune) []rune {
	res := make([]rune, len(s))
	idx := 0
	for i := len(s) - 1; i >= 0; i-- {
		res[idx] = s[i]
		idx++
	}
	return res
}

func main() {
	testCases := []string{
		"главрыба",
		"hello",
		"世界",
		"😊🐱",
		"a",
		"",
	}

	for _, test := range testCases {
		original := []rune(test)
		reversed := reverse(original)
		fmt.Printf("Original: %s -> Reversed: %s\n", test, string(reversed))
	}
}
