package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func flatten(root *TreeNode) {
    curr := root
    for curr != nil {
        if curr.Left != nil {
            prev := curr.Left
            for prev.Right != nil {
                prev = prev.Right
            }
            prev.Right = curr.Right
            curr.Right = curr.Left
            curr.Left = nil
        }
        curr = curr.Right
    }
}

// Helper function to print the flattened tree
func printFlattened(root *TreeNode) {
    for root != nil {
        fmt.Print(root.Val, " -> ")
        root = root.Right
    }
    fmt.Println("NULL")
}

func main() {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 5}
    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 4}
    root.Right.Right = &TreeNode{Val: 6}

    flatten(root)
    printFlattened(root)
}
