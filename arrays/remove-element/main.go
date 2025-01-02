package main

import "fmt"

func main() {
	nums1 := replaceElement2([]int{3, 4, 3, 3, 2, 2, 3}, 3)
	fmt.Println(nums1)
}

func replaceElement1(nums1 []int, val int) []int {
	// Create a new slice to hold the merged result
	nums3 := make([]int, 0, len(nums1))

	// Add non-zero elements from nums1
	for i := range nums1 {
		if nums1[i] != val {
			nums3 = append(nums3, nums1[i])
		}
	}
	for i := range nums1 {
		if nums1[i] == val {
			nums3 = append(nums3, nums1[i])
		}
	}

	return nums3
}

//leetcode solution
func replaceElement2(nums1 []int, val int) int {
	i := 0
	for _, v := range nums1 {
		if v != val {
			nums1[i] = v
			i++
		}
	}
	fmt.Println(nums1)
	return i
}
