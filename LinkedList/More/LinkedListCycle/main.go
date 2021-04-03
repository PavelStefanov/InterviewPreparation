package main

import "fmt"

func main() {
	i := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}

	res := hasCycle(&i)

	fmt.Println(res)
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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	f := head
	s := head
	for f.Next != nil {
		f = f.Next
		s = s.Next
		if f.Next != nil {
			f = f.Next
		} else {
			return false
		}

		if s == f {
			return true
		}
	}

	return false
}
