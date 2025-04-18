package main

// Code
// Testcase
// Test Result
// Test Result
// 53. Maximum Subarray
// Medium
// Topics
// Companies
// Given an integer array nums, find the subarray with the largest sum, and return its sum.

// Example 1:

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum 6.
// Example 2:

// Input: nums = [1]
// Output: 1
// Explanation: The subarray [1] has the largest sum 1.
// Example 3:

// Input: nums = [5,4,-1,7,8]
// Output: 23
// Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104

// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(maxSum, currSum)
	}
	return maxSum

}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

// func main() {
// 	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
// 	fmt.Println(maxSubArray([]int{1}))
// 	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
// }
