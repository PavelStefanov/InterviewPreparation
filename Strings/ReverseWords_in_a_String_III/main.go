package main

import (
	"fmt"
)

func main() {
	str := "Let's take LeetCode contest"
	r := reverseWords(str)
	fmt.Print(r)
}

func reverseWords(s string) string {
	b := []byte(s)
	
	l, r := 0, 0
	for ; r <= len(s); r++ {
		if r == len(s) || s[r] == ' ' {
			swapWord(b, l, r-1)
			l = r + 1
		}
	}

	return string(b)
}

func swapWord(s []byte, l int, r int) {
	for l < r {
		t := s[l]
		s[l] = s[r]
		s[r] = t

		l++
		r--
	}
}
