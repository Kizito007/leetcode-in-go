package main

import "fmt"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// countNodes counts the total nodes in a complete binary tree
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getLeftHeight(root)
	rightHeight := getRightHeight(root)

	// If left and right heights are equal, it's a perfect tree
	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1 // 2^height - 1
	}

	// Otherwise, count nodes recursively
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// Get left height of tree
func getLeftHeight(node *TreeNode) int {
	height := 0
	for node != nil {
		height++
		node = node.Left
	}
	return height
}

// Get right height of tree
func getRightHeight(node *TreeNode) int {
	height := 0
	for node != nil {
		height++
		node = node.Right
	}
	return height
}

// Helper function to create a test tree
func createTestTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
		},
	}
}

func main() {
	root := createTestTree()
	fmt.Println(countNodes(root)) // Output: 6
}
