package main

import "fmt"

func main() {
	nums := []int{1}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	jumps, currentJump, maxCanJump := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {

		if i+nums[i] > maxCanJump {
			maxCanJump = i + nums[i]
		}
		if i == currentJump {
			jumps, currentJump = jumps+1, maxCanJump

			if currentJump >= len(nums)-1 {
				return jumps
			}
		}
	}
	return 0
}
