package main

// 238. Product of Array Except Self
// Medium
// Topics
// Companies
// Hint
// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:

// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.

// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

func productExceptSelf(nums []int) []int {
	prefix := 1
	var result []int = make([]int, len(nums))
	for index, value := range nums {
		result[index] = prefix
		prefix = prefix * value
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * postfix
		postfix = postfix * nums[i]
	}
	return result
}

// func main() {
// 	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
// }

// func productExceptSelf(nums []int) []int {
// 	var result []int = make([]int, len(nums))
// 	prefix := 1
// 	for index, num := range nums {
// 		result[index] = prefix
// 		prefix *= num
// 	}
// 	postfix := 1

// 	for i := len(nums) - 1; i >= 0; i-- {
// 		result[i] = result[i] * postfix
// 		postfix *= nums[i]
// 	}
// 	return result
// }
