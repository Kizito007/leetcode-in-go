package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		for start < end && !isAlphanumeric(s[start]) {
			start++
		}
		for start < end && !isAlphanumeric(s[end]) {
			end--
		}

		if strings.ToLower(string(s[start])) != strings.ToLower(string(s[end])) {
			return false
		}
		start++
		end--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
