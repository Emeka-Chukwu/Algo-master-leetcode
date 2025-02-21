package main

// 2348. Number of Zero-Filled Subarrays
// Medium
// Topics
// Companies
// Hint
// Given an integer array nums, return the number of subarrays filled with 0.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:

// Input: nums = [1,3,0,0,2,0,0,4]
// Output: 6
// Explanation:
// There are 4 occurrences of [0] as a subarray.
// There are 2 occurrences of [0,0] as a subarray.
// There is no occurrence of a subarray with a size more than 2 filled with 0. Therefore, we return 6.
// Example 2:

// Input: nums = [0,0,0,2,0,0]
// Output: 9
// Explanation:
// There are 5 occurrences of [0] as a subarray.
// There are 3 occurrences of [0,0] as a subarray.
// There is 1 occurrence of [0,0,0] as a subarray.
// There is no occurrence of a subarray with a size more than 3 filled with 0. Therefore, we return 9.
// Example 3:

// Input: nums = [2,10,2019]
// Output: 0
// Explanation: There is no subarray filled with 0. Therefore, we return 0.

// Constraints:

// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109

func zeroFilledSubarray(nums []int) int64 {
	var subArray, count int
	for _, value := range nums {
		if value != 0 {
			count = 0
			continue
		}
		count += 1
		subArray += count
	}
	return int64(subArray)
}

func zeroFilledSubarray2(nums []int) int64 {
	var subArray, count int
	for _, value := range nums {
		if value == 0 {
			count += 1
			subArray += count
		} else {
			count = 0
		}
	}
	return int64(subArray)
}

// func main() {
// 	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
// 	fmt.Println(zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
// 	fmt.Println(zeroFilledSubarray([]int{2, 10, 2019}))
// 	// fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))

// }
