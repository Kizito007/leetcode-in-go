package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	// Split the path by "/"
	tokens := strings.Split(path, "/")

	// Use a stack to process the tokens
	var stack []string

	for _, token := range tokens {
		// Skip empty tokens and "."
		if token == "" || token == "." {
			continue
		}
		// If token is "..", pop from the stack (if possible)
		if token == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			// Otherwise, it's a valid directory name; push it onto the stack
			stack = append(stack, token)
		}
	}

	// Build the canonical path by joining the stack with "/"
	// Always start with a "/" since it's an absolute path.
	return "/" + strings.Join(stack, "/")
}

func main() {
	// Test cases
	fmt.Println(simplifyPath("/home/"))          // Output: "/home"
	fmt.Println(simplifyPath("/../"))            // Output: "/"
	fmt.Println(simplifyPath("/home//foo/"))     // Output: "/home/foo"
	fmt.Println(simplifyPath("/a/./b/../../c/")) // Output: "/c"
}
