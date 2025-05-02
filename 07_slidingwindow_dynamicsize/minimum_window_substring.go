package main

import (
	"fmt"
	"math"
)

// 76. Minimum Window Substring
// Hard
// Topics
// Companies
// Hint
// Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

// The testcases will be generated such that the answer is unique.

// Example 1:

// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
// Example 2:

// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.
// Example 3:

// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.

// Constraints:

// m == s.length
// n == t.length
// 1 <= m, n <= 105
// s and t consist of uppercase and lowercase English letters.

// Follow up: Could you find an algorithm that runs in O(m + n) time?

func minWindow(s string, t string) string {
	existingCharacters := map[rune]bool{}
	checkCharacters := map[rune]int{}
	freqcharacters := map[rune]int{}
	for _, char := range t {
		existingCharacters[char] = true
		freqcharacters[char]++
	}
	miniMumValue := math.MaxInt
	miniValue := ""
	miniString := ""
	counter := 0
	right := 0
	left := 0
	leftStarter := 0

	for right < len(s) {
		miniString += string(s[right])
		if existingCharacters[rune(s[right])] {
			checkCharacters[rune(s[right])] = 1
			if left == 0 && right > leftStarter {
				left = right
			}
			counter++
			isEqual := true
			if len(checkCharacters) == len(existingCharacters) {
				for _, char := range t {
					if freqcharacters[char] != checkCharacters[char] {
						right++
						isEqual = false
					}
				}
				if isEqual {
					if counter < miniMumValue {
						miniMumValue = counter
					}
					miniValue = miniString
					miniString = ""
					counter = 0
					checkCharacters = map[rune]int{}
					right = left
					leftStarter = left
					left = 0
				}

			} else {
				right++
			}

		} else {
			right++
		}
	}
	return miniValue
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABCN"))
}

//////modify this algorithm
