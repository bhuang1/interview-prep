package main

import (
	"fmt"
	s "github.com/golang-collections/collections/stack"
)

/*
https://leetcode.com/problems/binary-tree-inorder-traversal/
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]
Example 2:
Input: root = []
Output: []
Example 3:
Input: root = [1]
Output: [1]
Example 4:
Input: root = [1,2]
Output: [2,1]
Example 5:
Input: root = [1,null,2]
Output: [1,2]
*/

func main() {
	rl := &TreeNode{Val: 6}
	left := &TreeNode{Val: 1, Right: rl}
	right := &TreeNode{Val: 8}
	root := &TreeNode{Val: 5, Left: left, Right: right}
	fmt.Println(inorderTraversal(root))
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	lst := []int{}
	stack := s.New()
	for root != nil || stack.Peek() != nil {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		root = stack.Pop().(*TreeNode)
		lst = append(lst, root.Val)
		root = root.Right
	}
	return lst
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
