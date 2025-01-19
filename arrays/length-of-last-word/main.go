package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
	n := len(s)
	length := 0
	for n > 0 {
		// decrement the length to compare last index
		n--
		if s[n] != ' ' {
			length += 1
		} else if length > 0 { // last space case
			return length
		}
	}
	return length
}
