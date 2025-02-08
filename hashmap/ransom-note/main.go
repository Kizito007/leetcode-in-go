package main

import "fmt"

func main() {
	ransomNote, magazine := "ada", "mazda"

	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	letterTrackers := make(map[rune]int)

	for _, char := range magazine {
		letterTrackers[char]++
	}

	for _, char := range ransomNote {
		if letterTrackers[char] == 0 {
			return false
		}
		letterTrackers[char]--
	}
	return true
}
