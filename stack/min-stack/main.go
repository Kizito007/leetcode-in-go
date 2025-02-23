package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64}, // start with a very high value
	}
}

func (this *MinStack) Push(val int) {
	// Append the value to the main stack.
	this.stack = append(this.stack, val)

	// Compare the new value with the current minimum (top of minStack).
	currentMin := this.minStack[len(this.minStack)-1]
	if val < currentMin {
		this.minStack = append(this.minStack, val)
	} else {
		// Otherwise, push the current minimum again.
		this.minStack = append(this.minStack, currentMin)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		// Remove the top elements from both stacks.
		this.stack = this.stack[:len(this.stack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1 // or handle error appropriately
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return -1 // or handle error appropriately
	}
	// The top of minStack always holds the current minimum.
	return this.minStack[len(this.minStack)-1]
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("Minimum:", minStack.GetMin()) // Expected output: -3
	minStack.Pop()
	fmt.Println("Top:", minStack.Top())        // Expected output: 0
	fmt.Println("Minimum:", minStack.GetMin()) // Expected output: -2
}
