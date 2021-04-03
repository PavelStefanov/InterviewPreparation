package main

import "fmt"

func main() {
	input := createList([]int{1, 4, 3, 2, 5, 2})

	res := partition(input, 3)

	printList(res)
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	h := &ListNode{arr[0], nil}
	c := h
	for i := 1; i < len(arr); i++ {
		c.Next = &ListNode{arr[i], nil}
		c = c.Next
	}

	return h
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// ListNode test
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{0, head}

	l := dummy
	c := dummy

	// [3,1,2]
	// 3
	// 	[1,4,3,2,5,2]
	// 3
	for c != nil && c.Next != nil {
		if c.Next.Val < x {
			if c.Next != l.Next {
				v := c.Next.Val
				// remove
				if c.Next.Next == nil {
					c.Next = nil
				} else {
					c.Next.Val = c.Next.Next.Val
					c.Next.Next = c.Next.Next.Next
				}

				// insert
				l.Next = &ListNode{v, l.Next}
				l = l.Next
			} else {
				c = c.Next
				l = l.Next
			}

		} else {
			c = c.Next
		}
	}

	return dummy.Next
}
