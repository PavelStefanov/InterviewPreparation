package main

import "fmt"

func main() {
	i := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	res := deleteDuplicates(&i)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1
// 1 1
// 1 1 1 1
// 1 1 2
// 1 2 2
// 1 2 3
// 1 1 2 2 3 3
// 1 2 3 3
// 1 1 2 3 3
// 1 2 2 3 3 4
// [1,2,3,3,4,4,5]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}

	p := dummy
	c := head

	for c != nil {

		for c.Next != nil && p.Next.Val == c.Next.Val {
			c = c.Next
		}

		if p.Next == c {
			p = p.Next
		} else {
			p.Next = c.Next
		}

		c = c.Next
	}

	return dummy.Next
}
