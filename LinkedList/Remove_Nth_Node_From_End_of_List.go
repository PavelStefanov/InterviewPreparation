package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	c := 1
	s, f := head, head

	for f.Next != nil {
		s = s.Next
		f = f.Next
		c++
		if f.Next != nil {
			f = f.Next
			c++
		}
	}

	cur := head

	del := c - n

	if del == 0 {
		return head.Next
	}

	for i := 0; i < c-n-1; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	s, f := head, head

	for i := 0; i < n; i++ {
		f = f.Next
	}

	if f == nil {
		return head.Next
	}

	for f.Next != nil {
		s = s.Next
		f = f.Next
	}

	s.Next = s.Next.Next

	return head
}
