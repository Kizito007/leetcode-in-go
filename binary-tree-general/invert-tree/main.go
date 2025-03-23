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

// invertTree recursively swaps the left and right children of all nodes
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap left and right subtrees
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// Helper function to create a test tree
func createTree() *TreeNode {
	return &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}
}

// Helper function to print the tree in order
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTree(root.Left)
	fmt.Print(root.Val, " ")
	printTree(root.Right)
}

func main() {
	// Create a test tree
	root := createTree()

	// Print original tree (in-order traversal)
	fmt.Print("Original Tree: ")
	printTree(root)
	fmt.Println()

	// Invert the tree
	invertedRoot := invertTree(root)

	// Print inverted tree (in-order traversal)
	fmt.Print("Inverted Tree: ")
	printTree(invertedRoot)
	fmt.Println()
}
