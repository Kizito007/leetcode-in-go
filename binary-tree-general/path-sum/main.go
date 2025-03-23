package main

import "fmt"

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// hasPathSum function
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

// Helper function to create a tree for testing
func createTestTree() *TreeNode {
	return &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 1},
			},
		},
	}
}

func main() {
	root := createTestTree()
	targetSum := 22
	fmt.Println("Has Path Sum:", hasPathSum(root, targetSum)) // Expected output: true
}
