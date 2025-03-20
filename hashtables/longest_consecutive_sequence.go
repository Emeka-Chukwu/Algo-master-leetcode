package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 128. Longest Consecutive Sequence
// Medium
// Topics
// Companies
// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
// Example 3:

// Input: nums = [1,0,1,2]
// Output: 3

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109

// func longestConsecutive(nums []int) int {
//     numSet := make(map[int]bool)

//     for _, num := range nums {
//         numSet[num] = true
//     }

//     longestStreak := 0

//     for num := range numSet {
//         if _, found := numSet[num-1]; !found {
//             currentNum := num
//             currentStreak := 1

//             for {
//                 if _, exists := numSet[currentNum+1]; exists {
//                     currentNum++
//                     currentStreak++
//                 } else {
//                     break
//                 }
//             }

//             longestStreak = max(longestStreak, currentStreak)
//         }
//     }

//     return longestStreak
// }

// // not optimize
// func longestConsecutive(nums []int) int {
// 	var numsMap map[int]bool = make(map[int]bool)
// 	var maxNum = nums[0]
// 	var minNum = nums[0]
// 	var maxConsecutive int
// 	var maxCurrentConsecutive int = 1
// 	for _, num := range nums {
// 		numsMap[num] = true
// 		if num > maxNum {
// 			maxNum = num
// 		}
// 		if num < minNum {
// 			minNum = num
// 		}
// 	}
// 	for i := minNum; i < maxNum+1; i++ {
// 		if _, ok := numsMap[i+1]; ok {
// 			maxCurrentConsecutive++
// 		} else {
// 			if maxCurrentConsecutive > maxConsecutive {
// 				maxConsecutive = maxCurrentConsecutive
// 			}
// 			maxCurrentConsecutive = 1
// 		}
// 	}
// 	return maxConsecutive
// }

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))
	fmt.Println(longestConsecutive([]int{1, 0, 1, 2, 10000000000000}))
}

func longestConsecutive(nums []int) int {
	var numsMap map[int]bool = map[int]bool{}
	for _, num := range nums {
		numsMap[num] = true
	}
	var maxConsecutive int
	for _, num := range nums {
		var maxCurrentConsecutive int = 1
		currentNum := num
		for {
			if _, ok := numsMap[currentNum+1]; ok {
				maxCurrentConsecutive++
				currentNum++
			} else {
				break
			}
		}
		if maxCurrentConsecutive > maxConsecutive {
			maxConsecutive = maxCurrentConsecutive
		}

	}
	return maxConsecutive
}
