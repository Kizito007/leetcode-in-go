package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 0, 0},
		{0, 1, 1},
		{1, 1, 1},
		{1, 0, 0},
	}
	gameOfLife(matrix)
	fmt.Println(matrix)
}

func gameOfLife(board [][]int) {
	rows, cols := len(board), len(board[0])
	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}, // Right, Down, Left, Up
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1}, // Diagonals
	}

	// Apply rules with temporary markers
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			liveNeighbors := 0

			// Count live neighbors
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && abs(board[nr][nc]) == 1 {
					liveNeighbors++
				}
			}

			// Rule 1 & 3: Live cell dies
			if board[r][c] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				board[r][c] = -1 // Alive → Dead
			}

			// Rule 4: Dead cell becomes alive
			if board[r][c] == 0 && liveNeighbors == 3 {
				board[r][c] = 2 // Dead → Alive
			}
		}
	}

	// Step 2: Finalize the board
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] > 0 {
				board[r][c] = 1 // Alive
			} else {
				board[r][c] = 0 // Dead
			}
		}
	}
}

// Helper to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
