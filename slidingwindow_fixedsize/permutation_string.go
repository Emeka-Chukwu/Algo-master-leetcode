package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 567. Permutation in String
// Medium
// Topics
// Companies
// Hint
// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

// In other words, return true if one of s1's permutations is the substring of s2.

// Example 1:

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:

// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

// Constraints:

// 1 <= s1.length, s2.length <= 104
// s1 and s2 consist of lowercase English letters.

func checkInclusion(s1 string, s2 string) bool {
	total := 0
	sumCheck := 0
	for index, char := range s1 {
		total += int(char)
		dataLoop := []rune(string(s2[index]))
		sumCheck += int(dataLoop[0])
	}
	if sumCheck == total {
		return true
	}
	n := len(s1)
	for index := n; index < len(s2); index++ {
		data := []rune(string(s2[index-len(s1)]))
		dataLoop := []rune(string(s2[index]))
		sumCheck += int(dataLoop[0])
		sumCheck -= int(data[0])
		if sumCheck == total {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}

// Input: s1 = "ab", s2 = "eidbaooo"
