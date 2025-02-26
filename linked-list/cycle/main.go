package main

import "fmt"

// ListNode defines a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle uses Floyd's Tortoise and Hare algorithm to determine if the linked list has a cycle.
func hasCycle(head *ListNode) bool {
	// Initialize both pointers to the head of the list.
	slow, fast := head, head

	// Traverse the list.
	// fast and fast.Next must be non-nil to proceed; otherwise, we've reached the end (no cycle).
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move slow pointer one step.
		fast = fast.Next.Next // Move fast pointer two steps.

		// If the two pointers meet, a cycle exists.
		if slow == fast {
			return true
		}
	}
	// If we exit the loop, fast reached the end of the list; therefore, no cycle exists.
	return false
}

// Helper function to create a cycle for testing purposes.
func createCycle(head *ListNode, pos int) *ListNode {
	if pos < 0 {
		return head
	}
	cycleEntry := head
	for i := 0; i < pos; i++ {
		cycleEntry = cycleEntry.Next
	}
	// Find the tail of the list.
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	// Create the cycle.
	tail.Next = cycleEntry
	return head
}

func main() {
	// Create a linked list: 3 -> 2 -> 0 -> -4
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	// Test with a cycle: tail connects to node at position 1 (n2)
	headWithCycle := createCycle(n1, 1)
	fmt.Println("List with cycle:", hasCycle(headWithCycle)) // Expected output: true

	// Test without a cycle:
	n4.Next = nil                                    // Remove cycle
	fmt.Println("List without cycle:", hasCycle(n1)) // Expected output: false
}
