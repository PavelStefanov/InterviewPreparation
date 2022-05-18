package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// [1] 1
// [1 2 3] 1
// [1 2 3] 2
// [1 2 3] 3
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	s, f := head, head

	for i := 0; i < n; i++ {
		f = f.Next
	}

	if f == nil {
		return s.Next
	}

	for f.Next != nil {
		f = f.Next
		s = s.Next
	}

	s.Next = s.Next.Next

	return head
}
