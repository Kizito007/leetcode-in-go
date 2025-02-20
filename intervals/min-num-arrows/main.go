package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	// If there are no balloons, no arrows are needed.
	if len(points) == 0 {
		return 0
	}

	// Sort the intervals (balloons) by their end coordinate.
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	fmt.Println(points)
	// Initialize the number of arrows with one arrow for the first balloon.
	arrows := 1
	// Position the first arrow at the end of the first balloon.
	arrowPosition := points[0][1]

	// Iterate over the remaining balloons.
	for i := 1; i < len(points); i++ {
		// If the current balloon starts after the last arrow position,
		// it means the current arrow can't burst it.
		if points[i][0] > arrowPosition {
			arrows++                     // Shoot a new arrow.
			arrowPosition = points[i][1] // Place the arrow at the end of this balloon.
		}
	}

	return arrows
}

func main() {
	// Example usage:
	points := [][]int{
		{10, 16},
		{2, 8},
		{1, 6},
		{7, 12},
	}
	fmt.Println(findMinArrowShots(points)) // Expected output: 2
}
