package main

import (
	"strings"
)

// 392. Is Subsequence
// Easy
// Topics
// Companies
// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

// Example 1:

// Input: s = "abc", t = "ahbgdc"
// Output: true
// Example 2:

// Input: s = "axc", t = "ahbgdc"
// Output: false

// Constraints:

// 0 <= s.length <= 100
// 0 <= t.length <= 104
// s and t consist only of lowercase English letters.

// Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109, and you want to check one by one to see if t has its subsequence. In this scenario, how would you change your code?

func isSubsequence(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if !strings.Contains(t, string(s[i])) {
			return false
		}
	}
	return true
}

// func main() {
// 	t, s := "ahbgdc", "abc"
// 	fmt.Println(isSubsequence(s, t))
// 	fmt.Println(isSubsequence("axc", "ahbgdc"))
// 	fmt.Println()
// 	fmt.Println(isSubsequence2(s, t))
// 	fmt.Println(isSubsequence2("axc", "ahbgdc"))
// }

// ////// the matching doesn't need to be corderly, we can achieve 0(N) for both time and space complexity
func isSubsequence2(s string, t string) bool {
	availableCharacters := map[string]string{}
	for i := 0; i < len(t); i++ {
		availableCharacters[string(t[i])] = string(t[i])
	}
	for i := 0; i < len(s); i++ {
		if _, ok := availableCharacters[string(s[i])]; !ok {
			return false
		}
	}
	return true
}
