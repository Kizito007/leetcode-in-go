package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Root is always first element in preorder
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// Find root in inorder to split left and right subtree
	rootIndex := 0
	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
			break
		}
	}

	// Left subtree (inorder: 0 to rootIndex-1)
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])

	// Right subtree (inorder: rootIndex+1 to end)
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

// Helper function to print inorder traversal
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Print(root.Val, " ")
	inorderTraversal(root.Right)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)
	fmt.Print("Inorder Traversal of Constructed Tree: ")
	inorderTraversal(root)
}
