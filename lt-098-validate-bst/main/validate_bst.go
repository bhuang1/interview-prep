package main

import "fmt"
/*
https://leetcode.com/problems/validate-binary-search-tree/
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/

func main() {
	fmt.Println(generateTrees(5))
}

func generateTrees(n int) []*TreeNode {
	return genTree(1, n)
}

func genTree(start,end int) []*TreeNode {
	lst := []*TreeNode{}
	if start > end {
		return append(lst, nil)
	}
	if start == end {
		return append(lst, &TreeNode{Val: start})
	}
	for i := start; i <= end; i++ {
		leftList := genTree(start, i-1)
		rightList := genTree(i+1, end)
		for _, left := range leftList {
			for _, right := range rightList {
				root := TreeNode{Val: i}
				root.Left = left
				root.Right = right
				lst = append(lst, &root)
			}
		}
	}
	return lst
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
