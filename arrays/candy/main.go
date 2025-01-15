package main

import (
	"fmt"
)

func main() {
	ratings := []int{1, 2, 2}

	fmt.Println(candy(ratings))
}

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}
	fmt.Println(candies)

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
		fmt.Println(candies)
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	fmt.Println(candies)

	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}
