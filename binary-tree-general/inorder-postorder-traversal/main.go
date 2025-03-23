package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// Root is always the last element in postorder
	rootVal := postorder[len(postorder)-1]
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
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])

	// Right subtree (inorder: rootIndex+1 to end)
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])

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
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := buildTree(inorder, postorder)
	fmt.Print("Inorder Traversal of Constructed Tree: ")
	inorderTraversal(root)
}
