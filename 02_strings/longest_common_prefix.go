package main

import (
	"strings"
)

// 14. Longest Common Prefix
// Easy
// Topics
// Companies
// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters if it is non-empty.

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		found := false
		for !found {
			val := strings.Index(strs[i], prefix)
			found = val != -1
			if val == -1 {
				prefix = prefix[0 : len(prefix)-1]
			}
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// func main() {

// 	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
// 	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))

// }
