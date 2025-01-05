package main

import "fmt"

func main() {
	nums := []int{6, 6, 5, 6, 5, 5}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	majority, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
		}
		if nums[i] == majority {
			fmt.Println(count)

			count++
		} else {
			fmt.Println(count)
			count--
		}
	}
	return majority
}
