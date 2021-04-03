package main

func sumLists() {
	list1 := createList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	list2 := createList([]int{5, 6, 4})

	res := addTwoNumbersReverseOrder(list1, list2)

	res.print()
}

func addTwoNumbersReverseOrder(l1 *ListNode, l2 *ListNode) *ListNode {
	d := &ListNode{0, nil}

	c := d
	r := 0

	for l1 != nil || l2 != nil {
		s := 0

		if l1 != nil {
			s = s + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			s = s + l2.Val
			l2 = l2.Next
		}

		if r != 0 {
			s = s + r
			r = 0
		}

		if s >= 10 {
			v := s % 10
			r = s / 10
			c.Next = &ListNode{v, nil}
		} else {
			c.Next = &ListNode{s, nil}
		}

		c = c.Next
	}

	if r != 0 {
		c.Next = &ListNode{r, nil}
	}

	return d.Next
}

func addTwoNumbersForwardOrder(l1 *ListNode, l2 *ListNode) *ListNode {
	// var l1n *ListNode = nil
	// if l1.Next != nil {
	// 	l1n = l1.Next
	// } else {
	// 	l1n = l1
	// }

	// var l2n *ListNode = nil
	// if l2.Next != nil {
	// 	l2n = l2.Next
	// } else {
	// 	l2n = l2
	// }

	// s := addTwoNumbersForwardOrder(l1.Next, l2.Next)

	return nil
}
