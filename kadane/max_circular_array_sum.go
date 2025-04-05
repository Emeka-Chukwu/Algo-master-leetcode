package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 918. Maximum Sum Circular Subarray
// Medium
// Topics
// Companies
// Hint
// Given a circular integer array nums of length n, return the maximum possible sum of a non-empty subarray of nums.

// A circular array means the end of the array connects to the beginning of the array. Formally, the next element of nums[i] is nums[(i + 1) % n] and the previous element of nums[i] is nums[(i - 1 + n) % n].

// A subarray may only include each element of the fixed buffer nums at most once. Formally, for a subarray nums[i], nums[i + 1], ..., nums[j], there does not exist i <= k1, k2 <= j with k1 % n == k2 % n.

// Example 1:

// Input: nums = [1,-2,3,-2]
// Output: 3
// Explanation: Subarray [3] has maximum sum 3.
// Example 2:

// Input: nums = [5,-3,5]
// Output: 10
// Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10.
// Example 3:

// Input: nums = [-3,-2,-3]
// Output: -2
// Explanation: Subarray [-2] has maximum sum -2.

// Constraints:

// n == nums.length
// 1 <= n <= 3 * 104
// -3 * 104 <= nums[i] <= 3 * 104
// Accepted
// 326.3K
// Submissions
// 693.2K
// Acceptance Rate
// 47.1%
// Topics
// Companies
// Hint 1
// Hint 2

// Hint 3
// Discussion (86)

func maxSubarraySumCircular(nums []int) int {
	maxSum := nums[0]
	currMax := nums[0]
	total := nums[0]
	minSum := nums[0]
	currMin := nums[0]

	for i := 1; i < len(nums); i++ {
		total += nums[i]
		currMax = max2(currMax+nums[i], nums[i])
		currMin = min(currMax+nums[i], nums[i])

		maxSum = max2(currMax, maxSum)
		minSum = min(minSum, currMin)
	}

	if maxSum < 0 {
		return maxSum
	}

	return max2(maxSum, total-minSum)

}

func max2(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))
}
