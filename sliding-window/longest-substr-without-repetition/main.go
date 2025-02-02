package main

import "fmt"

func main() {
	s := "abcabcbb"

	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]bool)
	left, maxLength := 0, 0

	for right := 0; right < len(s); right++ {
		// If duplicate is found, move left forward
		for charMap[s[right]] {
			delete(charMap, s[left])
			left++
		}
		// Add new character and update maxLength
		charMap[s[right]] = true
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
