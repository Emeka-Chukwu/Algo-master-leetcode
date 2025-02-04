package main

// Code
// Testcase
// Test Result
// Test Result
// 169. Majority Element
// Easy
// Topics
// Companies
// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

// Constraints:

// n == nums.length
// 1 <= n <= 5 * 104
// -109 <= nums[i] <= 109

// Follow-up: Could you solve the problem in linear time and in O(1) space?

func majorityElement(nums []int) int {
	var maxValue int
	var keyMax int
	var freq map[int]int = map[int]int{}
	for _, value := range nums {
		freq[value] += 1
	}
	for key, value := range freq {
		if value > maxValue {
			maxValue = value
			keyMax = key
		}
	}
	return keyMax
}

// func main() {
// 	fmt.Println(majorityElement([]int{3, 2, 3}))
// 	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
// }
