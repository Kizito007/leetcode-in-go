package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}

	fmt.Println(jumpGame(nums))
}

func jumpGame(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	maxCanJump := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, "i")
		if i > maxCanJump {
			return false
		}
		if i+nums[i] > maxCanJump {
			maxCanJump = i + nums[i]
			fmt.Println(maxCanJump)
		}
	}
	return true
}
