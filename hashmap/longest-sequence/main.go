package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Step 1: Insert all numbers into a set
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0

	// Step 2: Iterate through the set
	for num := range numSet {
		// Only start counting if num is the beginning of a sequence.
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			// Count consecutive numbers
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			// Update maximum length if current streak is longer
			if currentStreak > maxLength {
				maxLength = currentStreak
			}
		}
	}

	return maxLength
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // Expected output: 4, since the consecutive sequence is 1, 2, 3, 4.
}
