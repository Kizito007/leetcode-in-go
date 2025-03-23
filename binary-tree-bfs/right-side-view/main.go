package main

import (
	"fmt"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rightSideView returns the rightmost node of each level in a binary tree
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var rightmost int

		// Traverse the current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			rightmost = node.Val // Update the rightmost value

			// Add child nodes for the next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Add the last node (rightmost) of the level to result
		result = append(result, rightmost)
	}

	return result
}

// Helper function to create a test tree
func createTestTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}
	return root
}

func main() {
	root := createTestTree()
	fmt.Println("Right Side View:", rightSideView(root)) // Output: [1, 3, 4]
}
