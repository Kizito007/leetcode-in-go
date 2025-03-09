package main

import "fmt"

// ListNode defines a node in the singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// partition rearranges the list so that nodes with values less than x come before nodes with values greater or equal to x.
func partition(head *ListNode, x int) *ListNode {
	// Create two dummy nodes to serve as the starting points for the two lists.
	dummyLess := &ListNode{}
	dummyGreater := &ListNode{}

	// Pointers for the two lists.
	less := dummyLess
	greater := dummyGreater

	// Traverse the original list.
	for head != nil {
		if head.Val < x {
			less.Next = head // Append to the less list.
			less = less.Next // Move the pointer forward.
		} else {
			greater.Next = head    // Append to the greater list.
			greater = greater.Next // Move the pointer forward.
		}
		head = head.Next // Proceed to the next node in the original list.
	}

	// Important: Terminate the greater list to avoid a potential cycle.
	greater.Next = nil

	// Connect the less list with the greater list.
	less.Next = dummyGreater.Next

	// The head of the partitioned list is the next node after dummyLess.
	return dummyLess.Next
}

// Helper function to print the linked list.
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// Create an example list: 1 -> 4 -> 3 -> 2 -> 5 -> 2, and x = 3.
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 2}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 2}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6

	fmt.Print("Original List: ")
	printList(n1)

	// Partition the list with x = 3.
	result := partition(n1, 3)
	fmt.Print("Partitioned List: ")
	printList(result)
}
