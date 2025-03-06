package main

import "fmt"

// ListNode defines a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd removes the nth node from the end of the list.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Step 1: Create a dummy node that points to head.
	dummy := &ListNode{Next: head}

	// Step 2: Initialize two pointers at the dummy.
	slow, fast := dummy, dummy

	// Step 3: Advance fast pointer n+1 steps ahead.
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Step 4: Move both pointers until fast reaches the end.
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Step 5: Skip the nth node from the end.
	slow.Next = slow.Next.Next

	// Step 6: Return the updated list.
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
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example: List: 1 -> 2 -> 3 -> 4 -> 5, remove the 2nd node from the end.
	// The expected output is: 1 -> 2 -> 3 -> 5
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original List: ")
	printList(head)

	newHead := removeNthFromEnd(head, 2)
	fmt.Print("Modified List: ")
	printList(newHead)
}
