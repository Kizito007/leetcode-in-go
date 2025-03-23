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

// isSymmetric checks if a tree is symmetric
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror checks if two trees are mirrors of each other
func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return (t1.Val == t2.Val) &&
		isMirror(t1.Left, t2.Right) &&
		isMirror(t1.Right, t2.Left)
}

// Helper function to create a symmetric tree
func createSymmetricTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
}

// Helper function to create an asymmetric tree
func createAsymmetricTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}
}

func main() {
	// Create symmetric and asymmetric trees
	symmetricTree := createSymmetricTree()
	asymmetricTree := createAsymmetricTree()

	// Check if trees are symmetric
	fmt.Println("Is Symmetric Tree:", isSymmetric(symmetricTree))   // true
	fmt.Println("Is Asymmetric Tree:", isSymmetric(asymmetricTree)) // false
}
