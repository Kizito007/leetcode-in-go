package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers adds two numbers represented as linked lists.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Dummy head to simplify result list operations.
	current := dummy     // Pointer to build the result list.
	carry := 0           // To keep track of the carry.

	// Loop until both lists are exhausted and no carry remains.
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry // Start with the carry.
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Calculate the new digit and update the carry.
		carry = sum / 10
		newDigit := sum % 10

		// Create a new node with the new digit and move current pointer.
		current.Next = &ListNode{Val: newDigit}
		current = current.Next
	}

	// The result is in dummy.Next.
	return dummy.Next
}

// Helper function to create a linked list from a slice.
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to print a linked list.
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example: Adding (2 -> 4 -> 3) and (5 -> 6 -> 4)
	// These represent the numbers 342 and 465, respectively.
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})

	result := addTwoNumbers(l1, l2)
	fmt.Print("Result: ")
	printList(result) // Expected output: 7 -> 0 -> 8 (representing 807)
}
