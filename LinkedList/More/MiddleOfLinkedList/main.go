package main

import "fmt"

func main() {
	i := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}

	res := middleNode(&i)

	fmt.Println(res.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 1
// 1 1 1
// 1 1 1 1
// 1 1 1 1 1
// 1 1 1 1 1 1
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	p := head

	for p.Next != nil {
		head = head.Next
		p = p.Next
		if p.Next != nil {
			p = p.Next
		}
	}

	return head
}
