package main

import (
	"fmt"
)

func main() {
	r := canConstruct("a", "b")
	fmt.Print(r)
}

func canConstruct(ransomNote string, magazine string) bool {
	chars := make([]int, 26)

	for i := 0; i < len(magazine); i++ {
		chars[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		if chars[ransomNote[i]-'a'] < 1 {
			return false
		}
		chars[magazine[i]-'a']--
	}

	return true
}
