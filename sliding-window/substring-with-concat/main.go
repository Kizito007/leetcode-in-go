package main

import "fmt"

func main() {
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring("barfoothefoobarman", words))
}

func findSubstring(s string, words []string) []int {
	wordFrequency := make(map[string]int)
	length := len(words[0])
	result := []int{}

	for _, word := range words {
		wordFrequency[word]++
	}

	for i := 0; i < len(s)-length*len(words)+1; i++ {
		seen := make(map[string]int)

		for j := 0; j < len(words); j++ {
			nextIndex := i + j*length
			word := s[nextIndex : nextIndex+length]

			if _, ok := wordFrequency[word]; !ok {
				break
			}

			seen[word]++

			if seen[word] > wordFrequency[word] {
				break
			}

			if j+1 == len(words) {
				result = append(result, i)
			}
		}
	}

	return result
}
