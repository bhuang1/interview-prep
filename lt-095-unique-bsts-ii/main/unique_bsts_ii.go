package main

import "fmt"
/*
https://leetcode.com/problems/unique-binary-search-trees-ii/
Given an integer n, return all the structurally unique BST's (binary search trees),
which has exactly n nodes of unique values from 1 to n. Return the answer in any order.

Example 1:
Input: n = 3
Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

Example 2:
Input: n = 1
Output: [[1]]
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
