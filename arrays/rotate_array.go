package main

// 189. Rotate Array
// Medium
// Topics
// Companies
// Hint
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

// Example 1:

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Example 2:

// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

// Constraints:

// 1 <= nums.length <= 105
// -231 <= nums[i] <= 231 - 1
// 0 <= k <= 105

// Follow up:

// Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
// Could you do it in-place with O(1) extra space?
// //the same as move to right
func rotate(nums []int, k int) []int {
	result := make([]int, len(nums))
	for index, value := range nums {
		newPos := (index + k) % len(nums)
		result[newPos] = value
	}
	return result
}

// ////move to the left
func rotateLeft(nums []int, k int) []int {
	result := make([]int, len(nums))
	for index, value := range nums {
		newPos := ((index - k) + len(nums)) % len(nums)
		result[newPos] = value
	}
	return result
}

// func main() {
// 	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 1))
// 	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 2))
// 	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))
// 	fmt.Println(rotate([]int{-1, -100, 3, 99}, 1))
// 	fmt.Println(rotate([]int{-1, -100, 3, 99}, 2))

// 	fmt.Println()
// 	fmt.Println(rotateLeft([]int{1, 2, 3, 4, 5, 6, 7}, 1))
// 	fmt.Println(rotateLeft([]int{1, 2, 3, 4, 5, 6, 7}, 2))
// 	fmt.Println(rotateLeft([]int{1, 2, 3, 4, 5, 6, 7}, 3))
// 	fmt.Println(rotateLeft([]int{-1, -100, 3, 99}, 1))
// 	fmt.Println(rotateLeft([]int{-1, -100, 3, 99}, 2))
// }
