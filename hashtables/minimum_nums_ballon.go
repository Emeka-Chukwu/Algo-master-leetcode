package main

// Code
// Testcase
// Test Result
// Test Result
// 1189. Maximum Number of Balloons
// Easy
// Topics
// Companies
// Hint
// Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

// You can use each character in text at most once. Return the maximum number of instances that can be formed.

// Example 1:

// Input: text = "nlaebolko"
// Output: 1
// Example 2:

// Input: text = "loonbalxballpoon"
// Output: 2
// Example 3:

// Input: text = "leetcode"
// Output: 0

// Constraints:

// 1 <= text.length <= 104
// text consists of lower case English letters only.

// Note: This question is the same as 2287: Rearrange Characters to Make Target String.

// Accepted
// 262.1K
// Submissions
// 438.8K
// Acceptance Rate
// 59.7%
// Topics
// Companies
// Hint 1
// Hint 2
// Discussion (24)

func maxNumberOfBalloons(text string) int {
	freq := make(map[rune]int)

	for _, value := range text {
		freq[value]++
	}
	ballonCount := map[rune]int{'b': 1, 'a': 1, 'l': 2, 'o': 2, 'n': 1}
	minimumBallonCount := len(text)
	for key, count := range ballonCount {
		if freq[key] < count {
			return 0
		}
		totalContribution := freq[key] / count
		if totalContribution < minimumBallonCount {
			minimumBallonCount = totalContribution
		}
	}
	return minimumBallonCount
}

// func main() {
// 	fmt.Println(maxNumberOfBalloons2("nlaebbbolkllo"))
// 	fmt.Println(maxNumberOfBalloons2("loonbalxballpoon"))
// 	fmt.Println(maxNumberOfBalloons2("leetcod"))
// 	fmt.Println()
// 	fmt.Println(maxNumberOfBalloons("nlaebbbolkllo"))
// 	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
// 	fmt.Println(maxNumberOfBalloons("leetcod"))
// }

func maxNumberOfBalloons2(text string) int {
	freq := make(map[rune]int)
	for _, char := range text {
		freq[char]++
	}
	ballonCount := map[rune]int{'b': 1, 'a': 1, 'l': 2, 'o': 2, 'n': 1}

	minimumCBallonCount := len(text)

	for key, count := range ballonCount {
		if freq[key] < count {
			return 0
		}
		totalContribution := freq[key] / count
		if totalContribution < minimumCBallonCount {
			minimumCBallonCount = totalContribution
		}
	}
	return minimumCBallonCount
}
