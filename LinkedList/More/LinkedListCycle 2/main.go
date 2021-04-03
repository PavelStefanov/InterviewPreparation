package main

import "fmt"

func main() {
	i1 := &ListNode{3, &ListNode{2, &ListNode{0, &ListNode{-4, nil}}}}

	i1.Next.Next.Next = i1.Next

	res := detectCycle(i1)

	if res != nil {
		fmt.Println(res.Val)
	}
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
func detectCycle(head *ListNode) *ListNode {

	f := head
	s := head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next

		if s == f {
			s := head
			for s != f {
				f = f.Next
				s = s.Next
			}
			return s
		}
	}

	return nil
}
