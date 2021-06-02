package main

import (
	"fmt"
	"math"
)

/*
https://leetcode.com/problems/minimum-window-substring/
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every
character in t (including duplicates) is included in the window. If there is no such substring, return the
empty string "".

The testcases will be generated such that the answer is unique.
A substring is a contiguous sequence of characters within the string.
Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
*/

func main() {
	fmt.Println(minWindow("ADOBECODEBANC","ABC"))
}

func minWindow(s string, t string) string {
	m := make(map[rune]int)
	for _, char := range t {
		m[char] = m[char] + 1
	}

	begin := 0
	end := 0
	head := 0
	counter := len(m)
	curLen := math.MaxInt64
	for end < len(s) {
		char := rune(s[end])
		if val, exists := m[char]; exists {
			m[char] = val-1
			if m[char] == 0 {
				counter--
			}
		}
		end++
		for counter == 0 {
			curChar := rune(s[begin])
			if val, exists := m[curChar]; exists {
				m[curChar] = val + 1
				if m[curChar] > 0 {
					counter++
				}
			}
			if end-begin < curLen {
				curLen = end-begin
				head = begin
			}
			begin++
		}
	}
	if curLen == math.MaxInt64 {
		return ""
	}
	return s[head:head+curLen]

}