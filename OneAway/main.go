package main

import (
	"fmt"
)

type test struct {
	s string
	t string
	r bool
}

func main() {
	tests := []test{
		// test{
		// 	s: "a",
		// 	t: "aaaa",
		// 	r: false,
		// },
		// test{
		// 	s: "abc",
		// 	t: "abc",
		// 	r: false,
		// },
		test{
			s: "acd",
			t: "abcd",
			r: true,
		},
		test{
			s: "abcd",
			t: "acd",
			r: true,
		},
		test{
			s: "abad",
			t: "abcd",
			r: true,
		},
		test{
			s: "abc",
			t: "abcdf",
			r: false,
		},
		test{
			s: "abcdf",
			t: "abc",
			r: false,
		},
		test{
			s: "bcde",
			t: "abcde",
			r: true,
		},
		test{
			s: "abcde",
			t: "bcde",
			r: true,
		},
		test{
			s: "",
			t: "a",
			r: true,
		},
	}

	for _, t := range tests {
		res := isOneEditDistance(t.s, t.t)

		fmt.Println(t.r == res)
	}
}

func isOneEditDistance(s string, t string) bool {
	ls := len(s)
	lt := len(t)

	if ls-lt > 1 || ls-lt < -1 {
		return false
	}

	c, i, j := 0, 0, 0

	for i < ls && j < lt {
		if s[i] != t[j] {
			c++
			if c > 1 {
				return false
			}

			if ls == lt {
				i++
				j++
			} else if ls > lt {
				i++
			} else {
				j++
			}
		} else {
			i++
			j++
		}
	}

	if i < ls || j < lt {
		c++
	}

	return c == 1
}
