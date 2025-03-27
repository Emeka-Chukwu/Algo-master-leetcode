package main

// Code
// Testcase
// Test Result
// Test Result
// 974. Subarray Sums Divisible by K
// Medium
// Topics
// Companies
// Given an integer array nums and an integer k, return the number of non-empty subarrays that have a sum divisible by k.
// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [4,5,0,-2,-3,1], k = 5
// Output: 7
// Explanation: There are 7 subarrays with a sum divisible by k = 5:
// [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
// Example 2:

// Input: nums = [5], k = 9
// Output: 0

// Constraints:

// 1 <= nums.length <= 3 * 104
// -104 <= nums[i] <= 104
// 2 <= k <= 104
// Accepted
// 393.5K
// Submissions
// 708.5K
// Acceptance Rate
// 55.5%
// Topics
// Companies
// Similar Questions
// Discussion (187)

func subarraysDivByK(nums []int, k int) int {
	count := 0
	sum := 0
	counterMap := map[int]int{}
	counterMap[0] = 1

	for _, num := range nums {
		sum += num
		remainder := sum % k
		if remainder < 0 {
			remainder += k
		}
		if value, found := counterMap[remainder]; found {
			count += value
		}
		counterMap[remainder] += 1
	}
	return count
}

// func main() {
// 	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
// 	fmt.Println(subarraysDivByK([]int{5}, 9))
// }
