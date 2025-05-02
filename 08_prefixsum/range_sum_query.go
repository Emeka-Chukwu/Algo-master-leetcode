package main

// 303. Range Sum Query - Immutable
// Easy
// Topics
// Companies
// Given an integer array nums, handle multiple queries of the following type:

// Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
// Implement the NumArray class:

// NumArray(int[] nums) Initializes the object with the integer array nums.
// int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

// Example 1:

// Input
// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// Output
// [null, 1, -1, -3]

// Explanation
// NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
// numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
// numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
// numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3

// Constraints:

// 1 <= nums.length <= 104
// -105 <= nums[i] <= 105
// 0 <= left <= right < nums.length
// At most 104 calls will be made to sumRange.

// brute force approach; the because the for loops everytime the sumrange is invoked
type NumArray struct {
	nums []int
}

// func Constructor(nums []int) NumArray {
// 	return NumArray{nums: nums}
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	var total int
// 	for i := left; i <= right; i++ {
// 		total += this.nums[i]
// 	}
// 	return total
// }

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

// func main() {
// 	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
// 	fmt.Println(obj.SumRange(0, 2))
// 	fmt.Println(obj.SumRange(2, 5))
// 	fmt.Println(obj.SumRange(0, 5))
// }

func Constructor(nums []int) NumArray {
	var prefixSum []int = make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	return NumArray{nums: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right+1] - this.nums[left]
}
