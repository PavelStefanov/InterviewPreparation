package main

import "fmt"

func main() {
	i := ListNode{1, &ListNode{1, &ListNode{1, nil}}}

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

func deleteDuplicates(head *ListNode) *ListNode {
	c := head

	for c != nil && c.Next != nil {

		d := false
		s := head
		for s != nil {
			if c.Next.Val == s.Val && c.Next != s {
				d = true
				break
			}
			s = s.Next
		}
		if d == true {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}

	return head
}
