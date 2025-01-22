package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	result := convert(s, numRows)
	fmt.Println("Converted Zigzag:", result)
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// Create a slice of strings to hold each row
	rows := make([]string, numRows)
	curRow := 0
	goingDown := false

	for _, char := range s {
		rows[curRow] += string(char)

		// Change direction when reaching the top or bottom row
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	// Join all rows to form the final string
	return strings.Join(rows, "")
}
