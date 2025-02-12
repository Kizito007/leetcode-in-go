package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	// This map's keys are the sorted version of a string,
	// and its values are lists of strings (anagrams).
	anagramMap := make(map[string][]string)

	for _, s := range strs {
		// Convert string to a slice of characters.
		chars := strings.Split(s, "")
		// Sort the slice.
		sort.Strings(chars)
		// Create the key by joining the sorted characters.
		key := strings.Join(chars, "")

		// Append the original string to the list in the map.
		anagramMap[key] = append(anagramMap[key], s)
	}

	// Prepare the result by extracting all groups from the map.
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(input)
	fmt.Println(groups)
}
