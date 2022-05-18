package main

import "fmt"

func main() {
	r := strStr("mississippi", "issip")
	fmt.Print(r)
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	n := len(haystack)
	m := len(needle)

	if m > n {
		return -1
	}

	p := -1
	k := 0

	for i := 0; i < n; i++ {
		if k == 0 {
			if haystack[i] == needle[k] {
				p = i
				k++
			}
		} else {
			if haystack[i] != needle[k] {
				k = 0
				i = p

			} else {
				k++
			}
		}

		if k == len(needle) {
			return p
		}

	}

	p = -1
	return p
}
