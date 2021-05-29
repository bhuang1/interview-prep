package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	left := &TreeNode{Val: 3}
	right := &TreeNode{Val: 7}
	root := &TreeNode{Val: 5, Left: left, Right: right}

	fmt.Println(insertIntoBST(root, 4))
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val >= val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
