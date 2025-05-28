package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 19. Remove Nth Node From End of List
// Medium
// Topics
// Companies
// Hint
// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Example 1:

// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:

// Input: head = [1], n = 1
// Output: []
// Example 3:

// Input: head = [1,2], n = 1
// Output: [1]

// Constraints:

// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

// Follow up: Could you do this in one pass?

// Accepted
// 3.4M
// Submissions
// 7M
// Acceptance Rate
// 48.6%
// Topics
// Companies
// Hint 1
// Similar Questions
// Discussion (252)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodeLength = 1
	for a := head; a.Next != nil; a = a.Next {
		nodeLength++
	}
	var dummy *ListNode = &ListNode{Next: head}
	current := dummy
	for i := 0; i < nodeLength-n; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	return dummy.Next
}

// func main() {
// value := sliceToList([]int{1, 2, 3, 4, 5})
// 	// fmt.Println(removeNthFromEnd(value, 2))
// 	// fmt.Println("========== ============= ============= =========== ========= ==========")
// 	// fmt.Println(removeNthFromEnd1(value, 2))
// 	fmt.Println("========== ============= ============= =========== ========= ==========")
// 	fmt.Println(removeNthFromEnd3(value, 2))

// }

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	return &ListNode{
		Val:  nums[0],
		Next: sliceToList(nums[1:]),
	}
}

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func removeNthFromEnd1(head *ListNode, n int) *ListNode {

	var nodeLength = 1
	for a := head; a.Next != nil; a = a.Next {
		nodeLength += 1
	}
	var dummy *ListNode = &ListNode{Next: head}
	current := dummy
	for i := 0; i < nodeLength-n; i++ {

		current = current.Next
		fmt.Println(i, "2 method", current)
	}
	current.Next = current.Next.Next
	return dummy.Next

}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {

	// var nodeLength = 1
	// for a := head; a.Next != nil; a = a.Next {
	// 	nodeLength += 1
	// }

	// current := head
	// for i := 0; i < nodeLength-n; i++ {

	// 	current = current.Next
	// 	fmt.Println(i, "2 method", current)
	// }
	// current.Next = current.Next.Next
	// return head.Next
	var nodeLength = 1
	for a := head; a.Next != nil; a = a.Next {
		nodeLength += 1
	}

	current := head
	for i := 0; i < nodeLength-n; i++ {
		current = current.Next
		fmt.Println(i, "2 method", current)
	}
	current.Next = current.Next.Next
	return head

}

func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	currentIndex := 1
	var dummy *ListNode = &ListNode{Next: head}
	current := dummy

	for a := head; a.Next != nil; a = a.Next {

		// fmt.Println(current)
		// fmt.Println(current.Next)
		fmt.Println(currentIndex, "index")
		if currentIndex == 6-n {
			current.Next = current.Next.Next
			fmt.Println(current, "kkkkhhhh")
			fmt.Println(current.Next, "kkkkhhhhj")
			fmt.Println(current.Next.Val, "0000000")
			return current
		} else {
			current = current.Next
			fmt.Println(current, current.Val)
		}
		currentIndex++
	}
	fmt.Println(current, "kkkk")
	for a := current; a.Next != nil; a = a.Next {
		fmt.Println(a.Val, "llllllll")
	}

	return dummy.Next
}
