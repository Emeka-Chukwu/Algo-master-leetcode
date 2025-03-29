package main

import (
	"fmt"
	"math"
)

// 209. Minimum Size Subarray Sum
// Medium
// Topics
// Companies
// Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Example 1:

// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.
// Example 2:

// Input: target = 4, nums = [1,4,4]
// Output: 1
// Example 3:

// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

// Constraints:

// 1 <= target <= 109
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 104

// Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).

func minSubArrayLen(target int, nums []int) int {
	totalSum := 0
	minimumNumSum := math.MaxInt
	n := len(nums)
	left := 0
	right := 0
	for right < n {
		if totalSum < target {
			totalSum += nums[right]
			right++
		}
		if totalSum >= target {
			totalSum -= nums[left]
			left++
			if right-left < minimumNumSum && totalSum >= target {
				minimumNumSum = right - left
			}
			if right-left == 0 {
				return 1
			}
		}
	}
	if minimumNumSum == math.MaxInt {
		minimumNumSum = 0
	}
	return minimumNumSum
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1, 2, 6, 5, 1}))
}

// totalSum := 0
// 	left := 0
// 	minimumNum := math.MaxInt
// 	right := 0
// 	for right < len(nums) {
// 		if totalSum < target {
// 			totalSum += nums[right]
// 			right++
// 		}
// 		if totalSum >= target {
// 			totalSum -= nums[left]
// 			left++
// 			if (right-left) < minimumNum && totalSum >= target {
// 				minimumNum = right - left
// 			}
// 			if right-left == 0 {
// 				return 1
// 			}
// 		}

// 	}
// 	if minimumNum == math.MaxInt {
// 		return 0
// 	}

// 	return minimumNum
