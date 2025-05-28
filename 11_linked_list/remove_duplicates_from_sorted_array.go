package main

// 82. Remove Duplicates from Sorted List II
// Medium
// Topics
// premium lock icon
// Companies
// Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

// Example 1:

// Input: head = [1,2,3,3,4,4,5]
// Output: [1,2,5]
// Example 2:

// Input: head = [1,1,1,2,3]
// Output: [2,3]

// Constraints:

// The number of nodes in the list is in the range [0, 300].
// -100 <= Node.val <= 100
// The list is guaranteed to be sorted in ascending order.
// Accepted
// 917,467/1.8M
// Acceptance Rate
// 49.8%

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var checker map[int]int = make(map[int]int)
	current := head

	for current != nil {
		checker[current.Val] += 1
		current = current.Next
	}
	var dummy *ListNode = &ListNode{}
	tail := dummy
	current = head
	for current != nil {

		if value, _ := checker[current.Val]; value == 1 {
			tail.Next = &ListNode{Val: current.Val}
			tail = tail.Next

		}
		current = current.Next
	}

	return dummy.Next
}

// func main() {
// 	values := sliceToList1([]int{1, 2, 3, 3, 4, 4, 5})
// 	results := deleteDuplicates(values)
// 	fmt.Println(results)
// 	solutions(results)
// }

func sliceToList1(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	return &ListNode{
		Val:  nums[0],
		Next: sliceToList1(nums[1:]),
	}
}

func solutions(node *ListNode) {
	for node != nil {
		node = node.Next
	}
}
