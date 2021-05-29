package main

import "fmt"
/*
https://leetcode.com/problems/unique-binary-search-trees/

Given an integer n, return the number of structurally unique BST's (binary search trees)
which has exactly n nodes of unique values from 1 to n.

Example 1:
Input: n = 3
Output: 5

Example 2:
Input: n = 1
Output: 1
*/

func main() {
	fmt.Println(genNumTreesDP(5))
}

// DP solution!!
func genNumTreesDP(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	res[2] = 2
	for i := 3; i <= n; i++ {
		for j := 1; j <= i; j++ {
			res[i] += res[j-1] * res[i - j]
		}
	}
	return res[n]
}

func numTrees(n int) int {
	return genNumTreesRecur(1,n)
}

// Recursive solution to generating number of unique BSTs
func genNumTreesRecur(start, end int) int {
	if start >= end {
		return 1
	}
	count := 0
	for i := start; i <= end; i++ {
		leftCount := genNumTreesRecur(start, i-1)
		rightCount := genNumTreesRecur(i+1, end)
		count += leftCount * rightCount
	}
	return count
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}