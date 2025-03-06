package main

import "fmt"

// ListNode defines a node in the singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates removes all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list.
func deleteDuplicates(head *ListNode) *ListNode {
	// Create a dummy node that points to the head of the list.
	dummy := &ListNode{Next: head}
	// prev will track the last node before a group of duplicates.
	prev := dummy
	// cur is used to traverse the list.
	cur := head

	// Traverse the entire list.
	for cur != nil {
		// If the current node has a duplicate...
		if cur.Next != nil && cur.Val == cur.Next.Val {
			// Skip all nodes with the same value.
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			// Link prev to the node after the last duplicate.
			prev.Next = cur.Next
		} else {
			// No duplicates detected; move prev forward.
			prev = prev.Next
		}
		// Move cur forward in every iteration.
		cur = cur.Next
	}

	return dummy.Next
}

// Helper function to print the linked list.
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example: Create the list 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 3}
	n5 := &ListNode{Val: 4}
	n6 := &ListNode{Val: 4}
	n7 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7

	fmt.Print("Original List: ")
	printList(n1)

	// Remove duplicates.
	newHead := deleteDuplicates(n1)

	fmt.Print("Modified List: ")
	printList(newHead)
}
