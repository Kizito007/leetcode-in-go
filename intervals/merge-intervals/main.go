package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	result := merge(intervals)
	fmt.Println("Merged Intervals:")
	for _, interval := range result {
		fmt.Printf("[%d, %d] ", interval[0], interval[1])
	}
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		lastMerged := merged[len(merged)-1]
		current := intervals[i]

		if current[0] <= lastMerged[1] {
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}
