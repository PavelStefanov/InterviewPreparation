package main

import "fmt"

func palindrome() {
	list := createList([]int{5, 6, 4})

	res := isPalindrome(list)

	fmt.Print(res)
}

func isPalindrome(head *ListNode) bool {

	s := head
	f := head

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next

	}

	

	return false
}
