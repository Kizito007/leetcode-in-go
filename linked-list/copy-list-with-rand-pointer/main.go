package main

import "fmt"

// Node defines a linked list node with a random pointer.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList returns a deep copy of the list with random pointers.
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Step 1: Create a mapping from original nodes to their copies.
	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	// Step 2: Assign next and random pointers for the copies.
	cur = head
	for cur != nil {
		if cur.Next != nil {
			m[cur].Next = m[cur.Next]
		}
		if cur.Random != nil {
			m[cur].Random = m[cur.Random]
		}
		cur = cur.Next
	}

	// Return the head of the new list.
	return m[head]
}

// Helper function to print the list for demonstration.
func printList(head *Node) {
	for head != nil {
		randomVal := "nil"
		if head.Random != nil {
			randomVal = fmt.Sprintf("%d", head.Random.Val)
		}
		fmt.Printf("Val: %d, Random: %s -> ", head.Val, randomVal)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// Create a sample list:
	// Node 1 -> Node 2 -> Node 3
	// Where:
	// Node 1's random pointer points to Node 3.
	// Node 2's random pointer points to Node 1.
	// Node 3's random pointer points to Node 2.
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node1.Random = node3
	node2.Random = node1
	node3.Random = node2

	// Copy the list.
	copiedList := copyRandomList(node1)

	// Print the original and copied lists.
	fmt.Println("Original list:")
	printList(node1)
	fmt.Println("Copied list:")
	printList(copiedList)
}
