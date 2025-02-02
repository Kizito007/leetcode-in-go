package main

import "fmt"

func main() {
	s, t := "ADOBECODEBANC", "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	targetFreq := make(map[byte]int)
	for i := range t {
		targetFreq[t[i]]++
	}

	left, minLen, minStart, matched := 0, len(s)+1, 0, 0
	windowFreq := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		char := s[right]
		windowFreq[char]++

		if targetFreq[char] > 0 && windowFreq[char] <= targetFreq[char] {
			matched++
		}

		for matched == len(t) {
			if right-left+1 < minLen {
				minLen, minStart = right-left+1, left
			}

			windowFreq[s[left]]--
			if targetFreq[s[left]] > 0 && windowFreq[s[left]] < targetFreq[s[left]] {
				matched--
			}
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}
	return s[minStart : minStart+minLen]
}
