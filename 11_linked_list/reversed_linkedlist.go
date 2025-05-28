package main

import "fmt"

func reversedLinkedList(head *ListNode) *ListNode {

	current := head
	var previous *ListNode
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
		fmt.Println(previous)
	}
	fmt.Println(head)
	// for head.Next != nil {
	// 	fmt.Println(head.Next)
	// }
	return head
}

// func main() {
// 	values := sliceToList1([]int{1, 2, 3, 4, 5})
// 	reversedLinkedList(values)
// }

// func sliceToList1(nums []int) *ListNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}

// 	return &ListNode{
// 		Val:  nums[0],
// 		Next: sliceToList1(nums[1:]),
// 	}
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }
