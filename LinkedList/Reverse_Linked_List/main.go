package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head

		if next == nil {
			break
		}

		head = next
	}

	return head
}
