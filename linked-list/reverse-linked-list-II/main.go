package main

import "fmt"

// ListNode defines a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseBetween reverses the nodes of the linked list from position m to n.
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n {
		return head
	}

	// Create a dummy node to simplify edge cases.
	dummy := &ListNode{Next: head}
	pre := dummy

	// Move pre to the node just before position m.
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	// start will point to the first node of the sublist to reverse.
	start := pre.Next
	// cur is the node that will be reversed.
	cur := start.Next

	// Reverse the sublist from m to n.
	// We'll perform n-m iterations.
	for i := 0; i < n-m; i++ {
		start.Next = cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		cur = start.Next
	}

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
	// Example: Reverse the sublist from position 2 to 4.
	// List: 1 -> 2 -> 3 -> 4 -> 5, m = 2, n = 4.
	// Expected output: 1 -> 4 -> 3 -> 2 -> 5
	list := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original List: ")
	printList(list)

	result := reverseBetween(list, 2, 4)
	fmt.Print("Modified List: ")
	printList(result)
}
