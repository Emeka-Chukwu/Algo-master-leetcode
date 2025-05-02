package main

// Code
// Testcase
// Test Result
// Test Result
// 1512. Number of Good Pairs
// Easy
// Topics
// Companies
// Hint
// Given an array of integers nums, return the number of good pairs.

// A pair (i, j) is called good if nums[i] == nums[j] and i < j.

// Example 1:

// Input: nums = [1,2,3,1,1,3]
// Output: 4
// Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
// Example 2:

// Input: nums = [1,1,1,1]
// Output: 6
// Explanation: Each pair in the array are good.
// Example 3:

// Input: nums = [1,2,3]
// Output: 0

// Constraints:

// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100

func numIdenticalPairs(nums []int) int {
	countPairs := make(map[int]int)
	countedPair := 0

	for _, num := range nums {
		countedPair += countPairs[num]
		countPairs[num]++
	}
	return countedPair
}

// func main() {

// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println("numIdenticalPairs")
// 	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
// 	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))
// 	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
// }
