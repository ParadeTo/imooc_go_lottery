package main

import "fmt"

type TreeNode struct {
 	Val int
	Left *TreeNode
  Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	root := &TreeNode{}
	root.Val = 3
	left := &TreeNode{}
	left.Val = 9
	right := &TreeNode{}
	right.Val = 20
	root.Left = left
	root.Right = right
	fmt.Println(maxDepth(root))
}
