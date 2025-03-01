package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two sorted linked lists and returns the merged sorted list.
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Dummy node to help simplify edge cases.
	dummy := &ListNode{}
	current := dummy

	// Compare and merge the nodes from both lists.
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// Attach any remaining nodes from l1 or l2.
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	// The merged list starts from dummy.Next.
	return dummy.Next
}

// Helper function to print a linked list.
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Create first list: 1 -> 2 -> 4
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}

	// Create second list: 1 -> 3 -> 4
	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	// Merge the two lists.
	merged := mergeTwoLists(l1, l2)

	// Print the merged list.
	printList(merged) // Expected output: 1 1 2 3 4 4
}
