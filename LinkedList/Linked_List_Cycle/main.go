package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]bool)

	isCycle := false

	for head != nil {
		_, ok := m[head]
		if ok {
			isCycle = true
			break
		}

		m[head] = true
		head = head.Next
	}

	return isCycle
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		if head == fast {
			return true
		}

		head = head.Next
		fast = fast.Next.Next
	}

	return false
}
