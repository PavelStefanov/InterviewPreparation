package main

import "fmt"

func main() {
	i := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}

	res := findKthElement(&i, 6)

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
func findKthElement(head *ListNode, k int) *ListNode {
	c := 0

	for head != nil {
		if c == k {
			return head
		}

		if head.Next != nil {
			head = head.Next
			c++
		} else {
			break
		}

	}

	return nil
}
