package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if i < size-1 {
				node.Next = queue[0]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

// Helper function to print Next Pointers
func printNextPointers(root *Node) {
	for root != nil {
		temp := root
		for temp != nil {
			fmt.Print(temp.Val, " -> ")
			temp = temp.Next
		}
		fmt.Println("NULL")

		if root.Left != nil {
			root = root.Left
		} else {
			root = root.Right
		}
	}
}

func main() {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Right = &Node{Val: 7}

	connect(root)
	printNextPointers(root)
}
