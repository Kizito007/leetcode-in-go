package main

import "fmt"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor finds the LCA of two nodes in a binary tree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Base cases
	if root == nil || root == p || root == q {
		return root
	}

	// Search in left and right subtrees
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// If both left and right are not nil, root is the LCA
	if left != nil && right != nil {
		return root
	}

	// Else, return the non-nil child (LCA is in that subtree)
	if left != nil {
		return left
	}
	return right
}

// Helper function to create a test tree
func createTestTree() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	return root
}

func main() {
	root := createTestTree()
	p := root.Left  // Node 5
	q := root.Right // Node 1
	lca := lowestCommonAncestor(root, p, q)
	fmt.Println("LCA:", lca.Val) // Output: 3
}
