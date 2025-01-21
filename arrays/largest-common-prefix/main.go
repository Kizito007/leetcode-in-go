package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// Sort the strings by length
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	// Start with the first string as a reference
	prefix := strs[0]

	// Compare the prefix with each string in the array
	for _, str := range strs[1:] {
		for str[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1] // Shorten the prefix
		}
		// If the prefix becomes empty, there's no common prefix
		if prefix == "" {
			return ""
		}
	}

	return prefix
}
