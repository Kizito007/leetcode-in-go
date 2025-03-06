package main

import "fmt"

// ListNode defines a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup reverses the nodes of the list k at a time.
func reverseKGroup(head *ListNode, k int) *ListNode {
	// Check if there are at least k nodes left; if not, return head.
	cur := head
	for range k {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	// Reverse k nodes.
	var prev *ListNode = nil
	cur = head
	for range k {
		next := cur.Next // Store the next node.
		cur.Next = prev  // Reverse the current node's pointer.
		prev = cur       // Move prev to the current node.
		cur = next       // Advance cur to the next node.
	}

	// head now becomes the tail of the reversed group.
	// Recursively reverse the rest of the list and attach it.
	head.Next = reverseKGroup(cur, k)

	// prev is the new head of the reversed group.
	return prev
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
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example: Create a list 1->2->3->4->5->6->7->8 and set k=3.
	list := createList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Print("Original List: ")
	printList(list)

	result := reverseKGroup(list, 3)
	fmt.Print("Reversed in k-Group: ")
	printList(result)
}
