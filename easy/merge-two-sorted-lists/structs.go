package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(values []int) *ListNode {
	if values == nil {
		return nil
	}

	var header ListNode
	current := &header

	for _, value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}

	return header.Next
}

func (l *ListNode) ToArray() []int {
	var result []int
	c := l
	for c != nil {
		result = append(result, c.Val)
		c = c.Next
	}
	return result
}
