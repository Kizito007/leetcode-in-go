package main

import "fmt"

// TreeNode defines a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth computes the maximum depth of the binary tree.
func maxDepth(root *TreeNode) int {
	// Base case: if the node is nil, the depth is 0.
	if root == nil {
		return 0
	}

	// Recursively find the maximum depth of the left and right subtrees.
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// The depth of the current node is the larger of the two depths plus one.
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	// Example: Create a binary tree:
	//         1
	//        / \
	//       2   3
	//      / \
	//     4   5
	//
	// Maximum depth = 3 (1 -> 2 -> 4 or 1 -> 2 -> 5)

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println("Maximum Depth of Binary Tree:", maxDepth(root))
}
