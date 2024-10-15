package main

import (
	"fmt"
	"strings"
)

func AreCharsUnique(s string) bool {
	s = strings.ToLower(s)
	charSet := make(map[rune]bool)
	for _, char := range s {
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}
	return true
}

func main() {
	testString := "aabcd"
	fmt.Printf("%s : %t\n", testString, AreCharsUnique(testString))
}
