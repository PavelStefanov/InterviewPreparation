package main

import "fmt"

func main() {
	arr := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(arr)
	fmt.Print(arr)
}

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		t := s[l]
		s[l] = s[r]
		s[r] = t

		l++
		r--
	}
}
