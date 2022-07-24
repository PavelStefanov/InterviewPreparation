package main

import "fmt"

func main() {
	r := longestPalindrome("abccccdd")
	fmt.Println(r)
}

// "a"
// "aa"
// "aaa"
// "abcabc"
// "abcabcd"
func longestPalindrome(s string) int {
	result := 0
	m := make(map[byte]int)

	n := len(s)
	for i := 0; i < n; i++ {
		m[s[i]]++

		v, _ := m[s[i]]
		if v == 2 {
			result += 2
			m[s[i]] = 0
		}
	}

	for _, count := range m {
		if count > 0 {
			result++
			break
		}
	}

	return result
}
