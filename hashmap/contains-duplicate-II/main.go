package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := make(map[int]int)

	for i, n := range nums {
		if j, ok := numsMap[n]; ok {
			if i-j <= k {
				return true
			} else {
				numsMap[n] = i
			}
		} else {
			numsMap[n] = i
		}
	}

	return false
}

func main() {
	nums := []int{100, 2, 4, 200, 1, 3, 2}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
