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

// isSameTree checks if two binary trees are the same
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Helper function to create a test tree
func createTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
}

func main() {
	// Create two identical trees
	tree1 := createTree()
	tree2 := createTree()

	// Create a different tree
	tree3 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}

	// Test cases
	fmt.Println("Tree1 and Tree2 are same:", isSameTree(tree1, tree2)) // true
	fmt.Println("Tree1 and Tree3 are same:", isSameTree(tree1, tree3)) // false
}
