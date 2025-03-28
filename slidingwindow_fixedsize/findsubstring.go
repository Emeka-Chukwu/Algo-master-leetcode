package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 30. Substring with Concatenation of All Words
// Hard
// Topics
// Companies
// You are given a string s and an array of strings words. All the strings of words are of the same length.

// A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.

// For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
// Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.

// Example 1:

// Input: s = "barfoothefoobarman", words = ["foo","bar"]

// Output: [0,9]

// Explanation:

// The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
// The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.

// Example 2:

// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]

// Output: []

// Explanation:

// There is no concatenated substring.

// Example 3:

// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]

// Output: [6,9,12]

// Explanation:

// The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].
// The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].
// The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].

// Constraints:

// 1 <= s.length <= 104
// 1 <= words.length <= 5000
// 1 <= words[i].length <= 30
// s and words[i] consist of lowercase English letters.

func findSubstring(s string, words []string) []int {
	wordsInStrings := ""
	totalAsciiValue := 0
	checkWordTotal := 0
	result := make([]int, 0)
	for _, word := range words {
		wordsInStrings += word
	}
	n := len(wordsInStrings)
	for i := 0; i < n; i++ {
		totalAsciiValue += int(wordsInStrings[i])
	}
	for i := 0; i < n; i++ {
		checkWordTotal += int(s[i])
	}
	if totalAsciiValue == checkWordTotal {
		result = append(result, 0)
	}
	k := len(words[0])
	counter := 0
	for j := n; j < len(s); j += k {

		counter++
		step := 3
		for i := k * counter; ; i-- {

			checkWordTotal -= int(s[i-1])
			step--
			if step == 0 {
				break
			}
		}
		for i := j; i < j+k; i++ {
			checkWordTotal += int(s[i])
		}

		if totalAsciiValue == checkWordTotal {
			result = append(result, j+k-(k*len(words[0])))
		}
	}
	return result
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
}
