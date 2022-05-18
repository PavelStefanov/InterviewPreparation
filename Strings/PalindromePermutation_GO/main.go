package main

import "fmt"

func main() {
	input := "aabbhijkbbhijkbbhijkkjihijkkjihijkkjihijkkjih"

	res := findPalindromePermutation(input)

	fmt.Println(res)
}

func findPalindromePermutation(input string) bool {
	chars := make(map[byte]int)

	for i := 0; i < len(input); i++ {
		char := input[i]

		if val, ok := chars[char]; ok {
			chars[char] = val + 1
		} else {
			chars[char] = 1
		}
	}

	return isPalindrome(chars)
}

func isPalindrome(chars map[byte]int) bool {
	hasOneEven := false

	for _, count := range chars {
		remainderByTwo := count % 2
		if remainderByTwo != 0 {
			if hasOneEven {
				return false
			} else {
				hasOneEven = true
			}
		}
	}

	return true
}
