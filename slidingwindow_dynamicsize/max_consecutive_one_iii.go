package main

// Code
// Testcase
// Test Result
// Test Result
// 1004. Max Consecutive Ones III
// Medium
// Topics
// Companies
// Hint
// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

// Example 1:

// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
// Example 2:

// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
// Output: 10
// Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// Constraints:

// 1 <= nums.length <= 105
// nums[i] is either 0 or 1.
// 0 <= k <= nums.length
// Accepted
// 916.3K
// Submissions
// 1.4M
// Acceptance Rate
// 65.4%
// Topics
// Companies
// Hint 1
// Hint 2
// Hint 3
// Hint 4
// Similar Questions
// Discussion (98)

func longestOnes(nums []int, k int) int {
	freqMax := 0
	freqMap := map[int]int{}
	maxLength := 0
	left := 0

	for right, num := range nums {
		freqMap[num]++
		value := freqMap[num]
		if value > freqMax {
			freqMax = value
		}

		if right-left+1-freqMax > k {
			freqMap[nums[left]]--
			left++
		}

		if right-left+1 > freqMax {
			maxLength = right - left + 1
		}

	}
	return maxLength
}

// func main() {
// 	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
// 	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
// 	// fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
// }
