package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Calculate left products
	left := 1
	for i := 0; i < n; i++ {
		result[i] = left
		left *= nums[i]
		fmt.Println(result, result[i], left, "l")
	}

	// Multiply with right products
	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}
