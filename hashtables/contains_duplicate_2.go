package main

// Code
// Testcase
// Test Result
// Test Result
// 219. Contains Duplicate II
// Easy
// Topics
// Companies
// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

// Example 1:

// Input: nums = [1,2,3,1], k = 3
// Output: true
// Example 2:

// Input: nums = [1,0,1,1], k = 1
// Output: true
// Example 3:

// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false

// Constraints:

// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109
// 0 <= k <= 105

func containsNearbyDuplicate(nums []int, k int) bool {
	duplicateMap := map[int]int{}

	for index, num := range nums {
		if mapIndex, ok := duplicateMap[num]; ok {
			diff := index - mapIndex
			if diff <= k {
				return true
			}
		}
		duplicateMap[num] = index
	}
	return false
}

// func main() {
// 	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
// 	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
// 	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 1))
// }
