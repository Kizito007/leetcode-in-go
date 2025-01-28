package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// Calculate the area
		area := min(height[left], height[right]) * (right - left)

		// Update the maximum area
		if area > maxArea {
			maxArea = area
		}

		// Move the pointer of the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// Helper function to find the minimum of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))             // 16
	fmt.Println(maxArea([]int{1, 2, 1}))                   // 2
}
