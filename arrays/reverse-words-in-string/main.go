package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	words := strings.Fields(s)
	n := len(words)

	// Reverse the entire array
	reverse(words, 0, n-1)

	// Join the words back into a single string
	return strings.Join(words, " ")
}

func reverse(words []string, start, end int) {
	for start < end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}
}
