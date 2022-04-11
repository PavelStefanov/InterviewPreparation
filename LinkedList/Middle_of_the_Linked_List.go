package main

func middleNode(head *ListNode) *ListNode {
	s, f := head, head

	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	if f.Next == nil {
		return s
	} else {
		return s.Next
	}
}
