package main

import "fmt"

func isValid(s string) bool {
	// Create a stack using a slice of runes.
	stack := []rune{}

	// Loop over each character in the string.
	for _, char := range s {
		// If it's an opening bracket, push onto the stack.
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			// If it's a closing bracket, check if the stack is empty.
			if len(stack) == 0 {
				return false
			}

			// Pop the top element from the stack.
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Check if the popped bracket matches the current closing bracket.
			if (char == ')' && top != '(') ||
				(char == '}' && top != '{') ||
				(char == ']' && top != '[') {
				return false
			}
		}
	}

	// If the stack is empty, all brackets matched; otherwise, it's invalid.
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
