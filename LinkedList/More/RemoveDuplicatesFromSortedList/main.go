package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    c := head

	for head != nil && c.Next != nil {
		if c.Val == c.Next.Val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}

	return head
}
