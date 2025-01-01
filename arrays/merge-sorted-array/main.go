package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := merge3([]int{1, 9, 6, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

	fmt.Println(nums1)
}

func merge1(nums1 []int, m int, nums2 []int, n int) []int {
	// check number of array we are supposed to return; if m||n is 0 return empty
	// sort the array in a desc order
	var newNums1 []int
	var newNums2 []int

	for _, num := range nums1 {
		if num != 0 {
			newNums1 = append(newNums1, num)
		}
	}
	for _, num := range nums2 {
		if num != 0 {
			newNums2 = append(newNums2, num)
		}
	}

	sort.Ints(newNums1)
	sort.Ints(newNums2)
	nums3 := append(newNums1[:m], newNums2[:n]...)
	sort.Ints(nums3)

	// merge the number of array we sorted according to the int params
	return nums3
}

func merge2(nums1 []int, m int, nums2 []int, n int) []int {
	// Create a new slice to hold the merged result
	nums3 := make([]int, 0, m+n)

	// Add non-zero elements from nums1
	for i := 0; i < m; i++ {
		if nums1[i] != 0 {
			nums3 = append(nums3, nums1[i])
		}
	}

	// Add non-zero elements from nums2
	for i := 0; i < n; i++ {
		if nums2[i] != 0 {
			nums3 = append(nums3, nums2[i])
		}
	}

	// Sort the merged slice
	sort.Ints(nums3)

	return nums3
}

// works on leetcode
func merge3(nums1 []int, m int, nums2 []int, n int) []int {
	// Start from the end of nums1 and nums2
	i, j, k := m-1, n-1, m+n-1

	// Merge nums2 into nums1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			// Move the larger element to nums1[k]
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// If there are remaining elements in nums2, copy them
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
	return nums1
}
