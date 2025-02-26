package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	// Stack to store integers
	stack := []int{}

	// Iterate through each token
	for _, token := range tokens {
		switch token {
		// If the token is an operator, pop the top two numbers
		case "+":
			// Pop two numbers
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// Push the result of a + b
			stack = append(stack, a+b)
		case "-":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// Note: subtraction order matters (a - b)
			stack = append(stack, a-b)
		case "*":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, a*b)
		case "/":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// Integer division truncates toward zero
			stack = append(stack, a/b)
		default:
			// If the token is a number, convert it and push onto the stack
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(fmt.Sprintf("Invalid token %s", token))
			}
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func main() {
	tokens1 := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens1)) // Output: 9

	tokens2 := []string{"4", "13", "5", "/", "+"}
	// Explanation: 4 + (13 / 5) = 4 + 2 = 6 (since integer division of 13/5 yields 2)
	fmt.Println(evalRPN(tokens2)) // Output: 6
}
