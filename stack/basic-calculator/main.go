package main

import (
	"fmt"
	"unicode"
)

func calculate(s string) int {
	// Stack to store the result and sign before a parenthesis.
	stack := []int{}
	result := 0 // This will accumulate the final result.
	sign := 1   // This represents the current sign (1 for positive, -1 for negative).

	// Loop through each character in the string.
	for i := 0; i < len(s); i++ {
		ch := s[i]

		// If the character is a digit, build the full number.
		if unicode.IsDigit(rune(ch)) {
			num := 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			// Add the built number to the result, multiplied by the current sign.
			result += sign * num
			i-- // Adjust index since the outer loop will increment it.
		} else if ch == '+' {
			// '+' sets the sign to positive.
			sign = 1
		} else if ch == '-' {
			// '-' sets the sign to negative.
			sign = -1
		} else if ch == '(' {
			// When encountering an '(', push the current result and sign onto the stack.
			stack = append(stack, result) // Save the current result.
			stack = append(stack, sign)   // Save the current sign.
			// Reset result and sign for the new sub-expression.
			result = 0
			sign = 1
		} else if ch == ')' {
			// When encountering ')', pop the sign and previous result from the stack.
			// The sub-expression result inside the parentheses is already computed in 'result'.
			n := len(stack)
			prevSign := stack[n-1] // This is the sign before the '('.
			stack = stack[:n-1]
			prevResult := stack[len(stack)-1] // This is the result before the '('.
			stack = stack[:len(stack)-1]
			// Combine the sub-expression with the previous result.
			result = prevResult + prevSign*result
		}
		// Spaces and any other characters are ignored.
	}

	return result
}

func main() {
	expression1 := "1 + 1"
	expression2 := " 2-1 + 2 "
	expression3 := "(1+(4+5+2)-3)+(6+8)"

	fmt.Println(calculate(expression1)) // Expected output: 2
	fmt.Println(calculate(expression2)) // Expected output: 3
	fmt.Println(calculate(expression3)) // Expected output: 23
}
