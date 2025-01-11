package main

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{3, 1, 0, 5, 4}

	fmt.Println(hIndex(citations))
}

func hIndex(citations []int) int {
	// sort the array in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	h := 0
	for i, c := range citations {
		if c >= i+1 {
			h = i + 1
		} else {
			break
		}
	}
	return h
}
