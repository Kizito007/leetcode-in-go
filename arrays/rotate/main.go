package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	rotate(nums, 3)

	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // Handle cases where k > len(nums)

	// Reverse the entire array
	reverse(nums, 0, n-1)
	// Reverse the first k elements
	reverse(nums, 0, k-1)
	// Reverse the remaining elements
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	fmt.Println(nums)
}
