package main

import "fmt"

/*
https://leetcode.com/problems/find-all-anagrams-in-a-string/
Given two strings s and p, return an array of all the start indices of p's anagrams in s.
You may return the answer in any order.

Example 1:
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
 */
func main() {
	res := findAnagrams("cbaebabacd", "abc")
	fmt.Println(res)
}

func findAnagrams(s string, p string) []int {
	m := make(map[string]int)
	for _, c := range p {
		char := string(c)
		m[char] = m[char] + 1
	}
	counter := len(m)
	begin := 0
	end := 0

	res := []int{}
	for end < len(s) {
		endChar := string(s[end])
		if val, exists := m[endChar]; exists {
			m[endChar] = val - 1
			if m[endChar] == 0 {
				counter--
			}

		}
		end++
		for counter == 0 {
			sChar := string(s[begin])
			if val, exists := m[sChar]; exists {
				m[sChar] = val + 1
				if m[sChar] > 0 {
					counter++
				}
			}
			if end-begin == len(p) {
				res = append(res, begin)
			}
			begin++
		}
	}
	return res
}