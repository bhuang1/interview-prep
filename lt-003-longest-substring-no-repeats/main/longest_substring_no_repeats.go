package main

import "fmt"

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/
Given a string s, find the length of the longest substring without repeating characters.
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func main() {
	res := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	begin := 0
	end := 0
	counter := 0
	maxLen := 0
	for end < len(s) {
		endChar := string(s[end])
		if val, exists := m[endChar]; exists && val >= 1 {
			counter++
		}
		m[endChar] = m[endChar] + 1
		end++
		for counter > 0 { // making this a valid condition again
			beginChar := string(s[begin])
			ct := m[beginChar]
			if ct > 1 {
				counter--
			}
			m[beginChar] = m[beginChar] - 1
			begin++
		}
		if end-begin > maxLen {
			maxLen = end-begin
		}
	}
	return maxLen
}