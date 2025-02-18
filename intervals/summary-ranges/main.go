package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	ans := []string{}
	n := len(nums)

	for i := 0; i < n; i++ {
		start := nums[i]

		for i < n-1 && nums[i+1] == nums[i]+1 {
			i++
		}

		if start == nums[i] {
			ans = append(ans, fmt.Sprintf("%d", nums[i]))
			continue
		}

		ans = append(ans, fmt.Sprintf("%d->%d", start, nums[i]))
	}

	return ans
}

func main() {
	// Example usage:
	nums := []int{0, 1, 2, 4, 5, 7}
	ranges := summaryRanges(nums)
	fmt.Println(ranges) // Output: ["0->2", "4->5", "7"]
}
