package main

import "fmt"

func main() {
	// r := lengthOfLongestSubstring("pwwkew")
	// r := lengthOfLongestSubstring("abcabcbb")
	r := lengthOfLongestSubstring("abba")
	fmt.Print(r)
}

func lengthOfLongestSubstring(s string) int {
	chars := []rune(s)
	m := make(map[rune]int)
	f, r := 0, 0

	for i := 0; i < len(chars); i++ {
		v, ok := m[chars[i]]
		if !ok {
			m[chars[i]] = i
		} else {
			for ; f < v; f++ {
				delete(m, chars[f])
			}
			f++
			m[chars[i]] = i
		}

		l := i - f + 1
		if r < l {
			r = l
		}
	}

	return r
}
