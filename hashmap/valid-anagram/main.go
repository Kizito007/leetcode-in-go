package main

import "fmt"

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sTracker := make(map[rune]int)

	for _, i := range s {
		sTracker[i]++
	}

	for _, j := range t {
		if sTracker[j] == 0 {
			return false
		}
		sTracker[j]--
	}
	return true
}
