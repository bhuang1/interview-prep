package main

import (
	"fmt"
	"math"
)

/*
https://leetcode.com/problems/validate-binary-search-tree/
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/

func main() {
	rl := &TreeNode{Val: 6}
	left := &TreeNode{Val: 1, Right: rl}
	right := &TreeNode{Val: 8}
	root := TreeNode{Val: 5, Left: left, Right: right}
	fmt.Println(isValidBST(&root))
}

func isValidBST(root *TreeNode) bool {
	//success, _, _ := valid(root, math.MaxInt64, math.MinInt64)
	success := valid2(root, math.MinInt64, math.MaxInt64)
	return success
}

func valid(root *TreeNode, curMin, curMax int) (bool, int, int) {
	if root == nil {
		return true, curMin, curMax
	}
	ok1, leftMin, leftMax := valid(root.Left, curMin, curMax)
	ok2, rightMin, rightMax := valid(root.Right, curMin, curMax)
	cond := ok1 && ok2 && root.Val > leftMax && root.Val < rightMin

	return cond, min(root.Val, min(leftMin, rightMin)), max(root.Val, max(leftMax, rightMax))
}

// a more optimized solution that I didn't figure out myself
func valid2(root *TreeNode, curMin, curMax int) bool {
	if root == nil {
		return true
	}

	if root.Val >= curMax || root.Val <= curMin {
		return false
	}
	leftCond := valid2(root.Left, curMin, root.Val)
	rightCond := valid2(root.Right, root.Val, curMax)
	return leftCond && rightCond
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
