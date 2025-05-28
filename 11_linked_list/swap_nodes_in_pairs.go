package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 24. Swap Nodes in Pairs
// Medium
// Topics
// premium lock icon
// Companies
// Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

// Example 1:

// Input: head = [1,2,3,4]

// Output: [2,1,4,3]

// Explanation:

// Example 2:

// Input: head = []

// Output: []

// Example 3:

// Input: head = [1]

// Output: [1]

// Example 4:

// Input: head = [1,2,3]

// Output: [2,1,3]

// Constraints:

// The number of nodes in the list is in the range [0, 100].
// 0 <= Node.val <= 100
// Accepted
// 1,626,494/2.4M
// Acceptance Rate
// 67.1%

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	var dummy *ListNode = &ListNode{Next: head}
	current = dummy

	for current.Next != nil && current.Next.Next != nil {
		next := current.Next.Next
		temp := current.Next
		fmt.Println(next, temp)
		current.Next.Next, current.Next = temp, next

		current = next

	}
	return dummy.Next

}

func solutions1(node *ListNode) {
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}

func main() {
	values := sliceToList2([]int{1, 2, 3, 4})
	results := swapPairs(values)
	fmt.Println(results, "hhh")
	solutions1(results)
}

func sliceToList2(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	return &ListNode{
		Val:  nums[0],
		Next: sliceToList2(nums[1:]),
	}
}
