package main

import (
	"fmt"
	"unicode"
)

func main() {
	r := isPalindrome("0P")
	fmt.Print(r)
}

//
//ab ba
//b a ba
func isPalindrome(s string) bool {
	chars := []rune(s)

	l, r := 0, len(chars)-1

	for l < r {
		if !unicode.IsLetter(chars[l]) && !unicode.IsDigit(chars[l]) {
			l++
			continue
		}
		if !unicode.IsLetter(chars[r]) && !unicode.IsDigit(chars[r]) {
			r--
			continue
		}

		if unicode.ToLower(chars[l]) != unicode.ToLower(chars[r]) {
			return false
		}

		l++
		r--
	}

	return true
}
