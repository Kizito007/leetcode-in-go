package main

import "fmt"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BSTIterator implements an iterator for a BST
type BSTIterator struct {
	stack []*TreeNode
}

// Constructor initializes the iterator
func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	iterator.leftmostInorder(root)
	return iterator
}

// Push all left nodes onto stack
func (this *BSTIterator) leftmostInorder(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

// Next returns the next smallest element
func (this *BSTIterator) Next() int {
	// Pop the top element from the stack
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	// If it has a right child, process its left subtree
	if top.Right != nil {
		this.leftmostInorder(top.Right)
	}

	return top.Val
}

// HasNext returns true if there are more elements
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

// Helper function to create a test tree
func createTestTree() *TreeNode {
	return &TreeNode{
		Val:  7,
		Left: &TreeNode{Val: 3},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20},
		},
	}
}

func main() {
	root := createTestTree()
	iterator := Constructor(root)

	fmt.Println(iterator.Next())    // 3
	fmt.Println(iterator.Next())    // 7
	fmt.Println(iterator.HasNext()) // true
	fmt.Println(iterator.Next())    // 9
	fmt.Println(iterator.HasNext()) // true
	fmt.Println(iterator.Next())    // 15
	fmt.Println(iterator.HasNext()) // true
	fmt.Println(iterator.Next())    // 20
	fmt.Println(iterator.HasNext()) // false
}
