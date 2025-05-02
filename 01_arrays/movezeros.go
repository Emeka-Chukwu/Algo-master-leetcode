package main

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:

// Input: nums = [0]

// Constraints:

// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

func moveZeroes(nums []int) []int {
	// using two ponters
	var pos, zeroIndex = 0, len(nums) - 1
	var newArray []int = make([]int, len(nums))
	for _, value := range nums {
		if value == 0 {
			newArray[zeroIndex] = value
			zeroIndex--
		} else {
			newArray[pos] = value
			pos++
		}
	}
	return newArray
}

// func main() {
// 	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
// 	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
// 	fmt.Println(moveZeroes([]int{0}))
// }
