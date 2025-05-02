package main

// 438. Find All Anagrams in a String
// Medium
// Topics
// Companies
// Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

// Example 1:

// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".
// Example 2:

// Input: s = "abab", p = "ab"
// Output: [0,1,2]
// Explanation:
// The substring with start index = 0 is "ab", which is an anagram of "ab".
// The substring with start index = 1 is "ba", which is an anagram of "ab".
// The substring with start index = 2 is "ab", which is an anagram of "ab".

// Constraints:

// 1 <= s.length, p.length <= 3 * 104
// s and p consist of lowercase English letters.

func findAnagrams(s string, p string) []int {
	total := 0
	result := make([]int, 0)
	for _, char := range p {
		total += int(char)
	}
	charSum := 0
	for index, char := range s {
		charSum += int(char)
		if index >= len(p) {
			data := []rune(string(s[index-len(p)]))
			value := data[0]
			charSum -= int(value)
			if total == charSum {
				result = append(result, index+1-len(p))
			}
		} else {
			if total == charSum {
				result = append(result, 0)
			}
		}

	}
	return result
}

// func main() {
// 	fmt.Println(findAnagrams("cbaebabacd", "abc"))
// 	fmt.Println(findAnagrams("abab", "ab"))
// 	fmt.Println(findAnagrams("cbaebabacdcbaebabacdcbaebabacd", "abc"))
// }
