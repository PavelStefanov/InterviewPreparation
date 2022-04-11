package main

import "fmt"

func main() {
	str := "R"
	r := balancedStringSplit(str)
	fmt.Print(r)
}

func balancedStringSplit(s string) int {
	chars := []rune(s)

	r, c := 0, 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == 'L' {
			c++
		} else {
			c--
		}

		if c == 0 {
			r++
		}
	}

	return r
}
