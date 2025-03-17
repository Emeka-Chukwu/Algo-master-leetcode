package main

// Code
// Testcase
// Test Result
// Test Result
// 383. Ransom Note
// Easy
// Topics
// Companies
// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

// Each letter in magazine can only be used once in ransomNote.

// Example 1:

// Input: ransomNote = "a", magazine = "b"
// Output: false
// Example 2:

// Input: ransomNote = "aa", magazine = "ab"
// Output: false
// Example 3:

// Input: ransomNote = "aa", magazine = "aab"
// Output: true

// Constraints:

// 1 <= ransomNote.length, magazine.length <= 105
// ransomNote and magazine consist of lowercase English letters.

func canConstruct(ransomNote string, magazine string) bool {
	charFrequency := make(map[rune]int)
	for _, char := range magazine {
		charFrequency[char]++
	}
	for _, char := range ransomNote {
		value, _ := charFrequency[char]
		if value == 0 {
			return false
		}
		charFrequency[char]--
	}
	return true
}

// func main() {
// 	fmt.Println(canConstruct("a", "b"))
// 	fmt.Println(canConstruct("aa", "ab"))
// 	fmt.Println(canConstruct("aa", "aaa"))
// }
