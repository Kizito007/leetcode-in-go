package main

import "fmt"

func main() {
	nums := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	i := 0
	// comparison starts from the second element: it doesn't necessarily delte the other duplicates
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i++
		}
	}
	fmt.Println(nums)
	// return the length of the new array after removing duplicates
	return len(nums[:i+1])
}
