package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2
//

//
// 1 3

// 1
// 1

// 1 3 5
// 5 6 7

// 3 3 3
// 3 3 3

// 1 6 8
// 2 7 9
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	d := ListNode{}
	cur := &d

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}

		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return d.Next
}
