package main

// Code
// Testcase
// Test Result
// Test Result
// 3. Longest Substring Without Repeating Characters
// Medium
// Topics
// Companies
// Hint
// Given a string s, find the length of the longest substring without duplicate characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Constraints:

// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
// Accepted
// 7.2M
// Submissions
// 19.7M
// Acceptance Rate
// 36.5%

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	lastIndex := 0
	elementCheck := map[int]bool{}
	for _, char := range s {
		_, found := elementCheck[int(char)]
		for found {
			delete(elementCheck, int(s[lastIndex]))
			_, found = elementCheck[int(char)]
			lastIndex++
		}
		elementCheck[int(char)] = true
		if len(elementCheck) > maxLength {
			maxLength = len(elementCheck)
		}
	}
	return maxLength
}

func lengthOfLongestSubstringUsingRune(s string) int {
	maxLength := 0
	lastIndex := 0
	elementCheck := map[rune]bool{}
	for _, char := range s {
		for elementCheck[char] {
			delete(elementCheck, rune(s[lastIndex]))
			lastIndex++
		}
		elementCheck[char] = true
		if len(elementCheck) > maxLength {
			maxLength = len(elementCheck)
		}
	}
	return maxLength
}

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
// 	fmt.Println(lengthOfLongestSubstring("bbbbb"))
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))
// }
