package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 659. Split Array into Consecutive Subsequences
// Medium
// Topics
// Companies
// You are given an integer array nums that is sorted in non-decreasing order.

// Determine if it is possible to split nums into one or more subsequences such that both of the following conditions are true:

// Each subsequence is a consecutive increasing sequence (i.e. each integer is exactly one more than the previous integer).
// All subsequences have a length of 3 or more.
// Return true if you can split nums according to the above conditions, or false otherwise.

// A subsequence of an array is a new array that is formed from the original array by deleting some (can be none) of the elements without disturbing the relative positions of the remaining elements. (i.e., [1,3,5] is a subsequence of [1,2,3,4,5] while [1,3,2] is not).

// Example 1:

// Input: nums = [1,2,3,3,4,5]
// Output: true
// Explanation: nums can be split into the following subsequences:
// [1,2,3,3,4,5] --> 1, 2, 3
// [1,2,3,3,4,5] --> 3, 4, 5
// Example 2:

// Input: nums = [1,2,3,3,4,4,5,5]
// Output: true
// Explanation: nums can be split into the following subsequences:
// [1,2,3,3,4,4,5,5] --> 1, 2, 3, 4, 5
// [1,2,3,3,4,4,5,5] --> 3, 4, 5
// Example 3:

// Input: nums = [1,2,3,4,4,5]
// Output: false
// Explanation: It is impossible to split nums into consecutive increasing subsequences of length 3 or more.

// Constraints:

// 1 <= nums.length <= 104
// -1000 <= nums[i] <= 1000
// nums is sorted in non-decreasing order.

func isPossible(nums []int) bool {
	numFreq := map[int]int{}
	for _, num := range nums {
		numFreq[num] += 1
	}
	var secondValue int = 0
	firstArray := []int{}
	secondArray := []int{}

	for _, num := range nums {
		if value, ok := numFreq[num]; ok {
			if value == 0 {
				continue
			}
			if value > 1 || secondValue == 0 {
				firstArray = append(firstArray, num)
			}
			if value == 1 && secondValue != 0 {
				secondArray = append(secondArray, num)
			}
			if value > 1 {
				secondValue = num
			}
			numFreq[num] -= 1
		}
	}
	fmt.Println(firstArray)
	fmt.Println(secondArray)
	return len(firstArray) >= 3 && len(secondArray) >= 3
}

// func main() {
// 	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
// 	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}))
// 	fmt.Println(isPossible([]int{1, 2, 3, 4, 4, 5}))
// }
