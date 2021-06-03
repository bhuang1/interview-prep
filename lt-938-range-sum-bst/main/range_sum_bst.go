package main

import "fmt"

/*
https://leetcode.com/problems/range-sum-of-bst/
Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a
value in the inclusive range [low, high].

Example 1:
Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
Output: 32
Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.
*/
func main() {
	rl := &TreeNode{Val: 6}
	left := &TreeNode{Val: 1, Right: rl}
	right := &TreeNode{Val: 8}
	root := &TreeNode{Val: 5, Left: left, Right: right}
	res := rangeSumBST(root, 3, 7)
	fmt.Println(res)
}
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) +
		rangeSumBST(root.Right, low, high)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
