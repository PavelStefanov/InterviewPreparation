package main

import (
	"fmt"
	"sort"
)

func main() {
	r := isAnagram("rat", "car")
	fmt.Print(r)
}

func isAnagram(s string, t string) bool {
	c := len(s)
	m := make(map[byte]int)

	for i := 0; i < c; i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		v, ok := m[t[i]]
		if !ok || v == 0 {
			return false
		}

		m[t[i]]--
		c--
	}

	return c == 0
}

func isAnagramSort(s string, t string) bool {
	n := len(s)
	m := len(t)
	if n != m {
		return false
	}

	sChars := []rune(s)
	sort.Slice(sChars, func(i, j int) bool {
		return sChars[i] < sChars[j]
	})

	tChars := []rune(t)
	sort.Slice(tChars, func(i, j int) bool {
		return tChars[i] < tChars[j]
	})

	for i := 0; i < n; i++ {
		if sChars[i] != tChars[i] {
			return false
		}
	}

	return true
}
