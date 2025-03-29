package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 424. Longest Repeating Character Replacement
// Medium
// Topics
// Companies
// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can get after performing the above operations.

// Example 1:

// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.
// Example 2:

// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
// There may exists other ways to achieve this answer too.

// Constraints:

// 1 <= s.length <= 105
// s consists of only uppercase English letters.
// 0 <= k <= s.length
// Accepted
// 1M
// Submissions
// 1.8M
// Acceptance Rate
// 56.6%
// Topics
// Companies
// Similar Questions
// Discussion (176)

func characterReplacement(s string, k int) int {
	maxLength := 0
	maxFreq := 0
	characterMaps := map[byte]int{}
	left := 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		characterMaps[char]++

		if characterMaps[char] > maxFreq {
			maxFreq = characterMaps[char]
		}

		if right-left+1-maxFreq > k {
			characterMaps[s[left]]--
			left++
		}

		if right-left+1 > maxFreq {
			maxLength = right - left + 1
		}
	}
	return maxLength
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))

}
