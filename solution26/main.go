package main

import (
	"fmt"
	"strings"
)

func uniqueCheck(temp string) bool {
	lowercaseString := strings.ToLower(temp)
	var myMap map[rune]struct{} = make(map[rune]struct{}) // inicialization of map to store info about existing some rune

	for _, el := range lowercaseString {
		if _, ok := myMap[el]; ok == true {
			return false // If wee see the rune that was seen before -> not unique
		}
		myMap[el] = struct{}{} // Takes 0 bytes of data.
	}
	return true // If we didn't see the rune that was seen before -> unique
}

func main() {
	var temp string
	fmt.Scan(&temp)
	fmt.Println(uniqueCheck(temp))
}
