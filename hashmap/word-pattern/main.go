package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern, s := "abba", "dog cat cat dog"

	fmt.Println(wordPattern(pattern, s))
}

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	patternToWord := make(map[byte]string)
	wordToPattern := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w := words[i]

		if val, exists := patternToWord[p]; exists {
			if val != w {
				return false
			}
		} else {
			patternToWord[p] = w
		}

		if val, exists := wordToPattern[w]; exists {
			if val != p {
				return false
			}
		} else {
			wordToPattern[w] = p
		}
	}

	return true
}
