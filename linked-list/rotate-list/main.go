package main

import "fmt"

// ListNode defines a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// rotateRight rotates the list to the right by k places.
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Step 1: Find the length of the list and the tail node.
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// Step 2: Make the list circular by connecting tail to head.
	tail.Next = head

	// Step 3: Compute the effective rotations needed.
	k = k % length
	// If k is 0 after modulo, the list remains unchanged.
	if k == 0 {
		tail.Next = nil
		return head
	}

	// Step 4: Find the new tail, which is (length - k - 1) nodes from the beginning.
	stepsToNewTail := length - k
	newTail := head
	for i := 0; i < stepsToNewTail-1; i++ {
		newTail = newTail.Next
	}

	// The new head is the node after the new tail.
	newHead := newTail.Next
	// Step 5: Break the circle.
	newTail.Next = nil

	return newHead
}

// Helper function to print the linked list.
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	// Rotate the list to the right by 2 places.
	rotated := rotateRight(n1, 2)
	fmt.Print("Rotated List: ")
	printList(rotated) // Expected output: 4 5 1 2 3
}
