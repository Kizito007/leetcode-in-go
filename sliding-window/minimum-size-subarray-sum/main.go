package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	left, sum := 0, 0
	minLength := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sum += nums[right] // Expand window by adding nums[right]

		// Shrink window while sum is >= target
		for sum >= target {
			// Update minimum length
			minLength = min(minLength, right-left+1)
			sum -= nums[left] // Shrink window from the left
			left++
		}
	}

	// If minLength was not updated, return 0
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

// Helper function to get minimum of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})) // Output: 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))          // Output: 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1}))   // Output: 0
}
