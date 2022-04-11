package main

import (
	"fmt"
)

func main() {
	r := checkInclusion("ab", "eidbaooo")
	fmt.Print(r)
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Chars := []rune(s1)
	s1Len := len(s1)
	m1 := make(map[rune]int)
	s2Chars := []rune(s2)
	m2 := make(map[rune]int)

	for i := 0; i < s1Len; i++ {
		m1[s1Chars[i]]++
	}

	for i := 0; i < len(s2Chars); i++ {
		if i < s1Len {
			m2[s2Chars[i]]++
		} else {
			m2[s2Chars[i-s1Len]]--
			m2[s2Chars[i]]++
		}

		if match(m1, m2) {
			return true
		}
	}

	return false
}

func match(m1 map[rune]int, m2 map[rune]int) bool {
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
