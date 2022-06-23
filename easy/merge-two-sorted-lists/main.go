package main

import "fmt"

func main() {
	fmt.Println("Input: list1 = [1,2,4], list2 = [1,3,4]")
	fmt.Printf("Output: %v\n\n", mergeTwoLists(NewList([]int{1, 2, 4}), NewList([]int{1, 3, 4})).ToArray())

	fmt.Println("Input: list1 = [], list2 = []")
	fmt.Printf("Output: %v\n\n", mergeTwoLists(NewList([]int{}), NewList([]int{})).ToArray())

	fmt.Println("Input: list1 = [], list2 = [0]")
	fmt.Printf("Output: %v\n\n", mergeTwoLists(NewList([]int{}), NewList([]int{0})).ToArray())

	fmt.Println("Input: list1 = [1,2,4], list2 = [1,3,4]")
	fmt.Printf("Output: %v\n\n", recursiveMergeToList(NewList([]int{1, 2, 4}), NewList([]int{1, 3, 4})).ToArray())

	fmt.Println("Input: list1 = [], list2 = []")
	fmt.Printf("Output: %v\n\n", recursiveMergeToList(NewList([]int{}), NewList([]int{})).ToArray())

	fmt.Println("Input: list1 = [], list2 = [0]")
	fmt.Printf("Output: %v\n\n", recursiveMergeToList(NewList([]int{}), NewList([]int{0})).ToArray())
}

/*
type ListNode struct {
	Val  int
	Next *ListNode
}
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode { // Time: O(n + m); Space: O(1)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	current := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}

	return head
}

func recursiveMergeToList(l1 *ListNode, l2 *ListNode) *ListNode { // Time: O(n + m); Space: O(n + m)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = recursiveMergeToList(l1.Next, l2)
		return l1
	}

	l2.Next = recursiveMergeToList(l1, l2.Next)
	return l2
}
