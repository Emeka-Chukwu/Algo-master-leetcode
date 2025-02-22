package main

import (
	"fmt"
	"strings"
)

// Code
// Testcase
// Test Result
// Test Result
// 125. Valid Palindrome
// Easy
// Topics
// Companies
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:

// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

// Constraints:

// 1 <= s.length <= 2 * 105
// s consists only of printable ASCII characters.

func isPalindrome(s string) bool {
	str := filterAlphaNumeric(s)
	left, right := 0, len(str)-1
	for left < right {
		if strings.ToLower(string(str[left])) != strings.ToLower(string(str[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func filterAlphaNumeric(str string) string {
	results := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		char := str[i]
		if (char >= '0' && char <= '9') || (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			results = append(results, char)
		}
	}
	return string(results)
}

func main() {

	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))

}
