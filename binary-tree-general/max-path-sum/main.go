package main

import (
	"fmt"
	"math"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxPathSum function to find the maximum path sum
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	calculateMaxPath(root, &maxSum)
	return maxSum
}

// Helper function to calculate the max path sum
func calculateMaxPath(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	// Compute max sum from left and right subtrees (ignore negative sums)
	leftSum := max(0, calculateMaxPath(node.Left, maxSum))
	rightSum := max(0, calculateMaxPath(node.Right, maxSum))

	// Compute max path sum **including** current node
	currentMax := node.Val + leftSum + rightSum

	// Update global max path sum
	*maxSum = max(*maxSum, currentMax)

	// Return max path sum **including only one child** (left or right)
	return node.Val + max(leftSum, rightSum)
}

// Helper function to get max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test function to create a sample tree
func createTestTree() *TreeNode {
	return &TreeNode{
		Val:  -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
}

func main() {
	root := createTestTree()
	fmt.Println("Maximum Path Sum:", maxPathSum(root)) // Expected Output: 42
}
