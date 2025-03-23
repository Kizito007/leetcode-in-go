package main

import "fmt"

// TreeNode structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sumNumbers function to calculate sum of root-to-leaf numbers
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

// Helper DFS function
func dfs(node *TreeNode, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum = currentSum*10 + node.Val

	// If it's a leaf node, return the accumulated sum
	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	// Recursively sum the left and right subtrees
	return dfs(node.Left, currentSum) + dfs(node.Right, currentSum)
}

// Helper function to create a sample test tree
func createTestTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
}

func main() {
	root := createTestTree()
	fmt.Println("Sum of Root-to-Leaf Numbers:", sumNumbers(root)) // Expected Output: 262 (124 + 125 + 13)
}
