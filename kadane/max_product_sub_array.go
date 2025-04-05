package main

import "fmt"

// 152. Maximum Product Subarray
// Medium
// Topics
// Companies
// Given an integer array nums, find a subarray that has the largest product, and return the product.

// The test cases are generated so that the answer will fit in a 32-bit integer.

// Example 1:

// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.
// Example 2:

// Input: nums = [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

// Constraints:

// 1 <= nums.length <= 2 * 104
// -10 <= nums[i] <= 10
// The product of any subarray of nums is guaranteed to fit in a 32-bit integer.
// Accepted
// 1.6M
// Submissions
// 4.6M
// Acceptance Rate
// 34.6%
// Topics
// Companies
// Similar Questions
// Discussion (266)

func maxProduct(nums []int) int {
	maxProductSubarray := nums[0]
	currProductSubarray := nums[0]

	for i := 1; i < len(nums); i++ {
		currProductSubarray = max3(nums[i], currProductSubarray*nums[i])
		maxProductSubarray = max3(currProductSubarray, maxProductSubarray)
	}
	return maxProductSubarray
}

func max3(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-3, -2, -3}))
}
