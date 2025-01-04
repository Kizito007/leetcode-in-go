package main

import "fmt"

func main() {
	nums := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	j := 2
	// comparison starts from the second element: it doesn't necessarily delte the other duplicates
	for k := 2; k < len(nums); k++ {
		if nums[k] != nums[j-2] {
			nums[j] = nums[k]
			j++
		}
	}
	fmt.Println(nums, j)
	// return the length of the new array after removing duplicates
	return len(nums[:j+1])
}
