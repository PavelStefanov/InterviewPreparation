package main

import "fmt"

// ListNode comment
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) print() {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func createList(arr []int) *ListNode {
	if arr == nil {
		return nil
	}

	dummy := &ListNode{0, nil}
	c := dummy

	for _, v := range arr {
		c.Next = &ListNode{v, nil}
		c = c.Next
	}

	return dummy.Next
}
