package main

import (
	"fmt"
)

func reverse(temp string) string {
	s := []byte(temp)

	i, j := 0, len(s)-1
	for i < j {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		i++
		j--
	}
	prev := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || i == len(s)-1 {
			l, r := prev, i-1

			if i == len(s)-1 {
				r = i
			}

			for l < r {
				temp := s[l]
				s[l] = s[r]
				s[r] = temp
				l++
				r--
			}
			prev = i + 1
		}
	}

	return string(s)
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverse(s))
}
