package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	result := fullJustify(words, maxWidth)

	for _, line := range result {
		fmt.Printf("[%s]\n", line)
	}
}

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	cur := []string{}
	num_of_letters := 0

	for _, word := range words {
		if len(word)+len(cur)+num_of_letters > maxWidth {
			for i := 0; i < maxWidth-num_of_letters; i++ {
				cur[i%max(1, len(cur)-1)] += " "
			}
			res = append(res, strings.Join(cur, ""))
			cur = cur[:0]
			num_of_letters = 0
		}
		cur = append(cur, word)
		num_of_letters += len(word)
	}

	lastLine := strings.Join(cur, " ")
	for len(lastLine) < maxWidth {
		lastLine += " "
	}
	res = append(res, lastLine)

	return res
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
