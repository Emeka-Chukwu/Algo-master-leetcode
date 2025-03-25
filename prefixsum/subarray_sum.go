package main

import "fmt"

// 560. Subarray Sum Equals K
// Medium
// Topics
// Companies
// Hint
// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:

// Input: nums = [1,1,1], k = 2
// Output: 2
// Example 2:

// Input: nums = [1,2,3], k = 3
// Output: 2

// Constraints:

// 1 <= nums.length <= 2 * 104
// -1000 <= nums[i] <= 1000
// -107 <= k <= 107

func subarraySum1(nums []int, k int) int {
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		currentSum := prefixSum[i] + nums[i]
		prefixSum[i+1] = currentSum
		if currentSum == k {
			return i + 1
		}

	}
	return 0
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{1, 2, 3, 3, 1, 1}, 4))

}

func subarraySum(nums []int, k int) int {
	count := 0
	sumMap := map[int]int{}
	sumMap[0] = 1
	sum := 0

	for _, num := range nums {
		sum += num
		if value, ok := sumMap[sum-k]; ok {
			count += value
		}
		sumMap[sum] += 1
	}
	return count
}
