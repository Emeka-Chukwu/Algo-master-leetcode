package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 1014. Best Sightseeing Pair
// Medium
// Topics
// Companies
// Hint
// You are given an integer array values where values[i] represents the value of the ith sightseeing spot. Two sightseeing spots i and j have a distance j - i between them.

// The score of a pair (i < j) of sightseeing spots is values[i] + values[j] + i - j: the sum of the values of the sightseeing spots, minus the distance between them.

// Return the maximum score of a pair of sightseeing spots.

// Example 1:

// Input: values = [8,1,5,2,6]
// Output: 11
// Explanation: i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11
// Example 2:

// Input: values = [1,2]
// Output: 2

// Constraints:

// 2 <= values.length <= 5 * 104
// 1 <= values[i] <= 1000
// Accepted
// 205.2K
// Submissions
// 327.7K
// Acceptance Rate
// 62.6%
// Topics
// Companies
// Hint 1
// Discussion (138)

func maxScoreSightseeingPair(values []int) int {
	maxScore := 0
	n := len(values)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			score := values[i] + values[j] + i - j
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
	fmt.Println(maxScoreSightseeingPair([]int{1, 2}))
	fmt.Println()
	fmt.Println(maxScoreSightseeingPair2([]int{8, 1, 5, 2, 6}))
	fmt.Println(maxScoreSightseeingPair2([]int{1, 2}))
}

// values[i] + values[j] + i - j
// //optimised
func maxScoreSightseeingPair2(values []int) int {
	maxScore := 0
	currMax := values[0]

	for i := 1; i < len(values); i++ {
		maxScore = max4(maxScore, currMax+values[i]-i)
		currMax = max4(currMax, values[i]+1)
	}
	return maxScore
}

func max4(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
