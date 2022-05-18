package main

import (
	"fmt"
	"unicode"
)

func main() {
	r := letterCasePermutation("a1b2c3")
	fmt.Print(r)
}

func letterCasePermutation(s string) []string {
	chars := []rune(s)
	result := []string{}

	dfs(0, chars, &result)

	return result
}

func dfs(start int, chars []rune, result *[]string) {
	if start < len(chars) {
		dfs(start+1, chars, result)
		if unicode.IsLetter(chars[start]) {
			isUpper := unicode.IsUpper(chars[start])
			if isUpper {
				chars[start] = unicode.ToLower(chars[start])
				dfs(start+1, chars, result)
				chars[start] = unicode.ToUpper(chars[start])
			} else {
				chars[start] = unicode.ToUpper(chars[start])
				dfs(start+1, chars, result)
				chars[start] = unicode.ToLower(chars[start])
			}
		}
	} else {
		*result = append(*result, string(chars))
	}
}
