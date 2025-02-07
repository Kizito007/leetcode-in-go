package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix)) // Output: [1 2 3 6 9 8 7 4 5]
}

func spiralOrder(matrix [][]int) []int {

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	var result []int

	for top <= bottom && left <= right {
		// Traverse from left to right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
			fmt.Println(result)
		}
		top++

		// Traverse from top to bottom
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Traverse from right to left
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Traverse from bottom to top
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}
