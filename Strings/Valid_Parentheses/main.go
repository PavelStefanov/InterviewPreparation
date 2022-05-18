package main

import "fmt"

// {
// {({
// {}
// {()}
// {([])}
// {)
// })]
func main() {
	r := isValid("{[]}")
	fmt.Print(r)
}

func isValid(s string) bool {
	chars := []rune(s)
	stack := []rune{}

	for i := 0; i < len(chars); i++ {
		if chars[i] == '(' || chars[i] == '{' || chars[i] == '[' {
			stack = append(stack, chars[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			n := len(stack)
			last := stack[n-1]
			stack = stack[:n-1]

			if (last == '(' && chars[i] == ')') || (last == '{' && chars[i] == '}') || (last == '[' && chars[i] == ']') {
				continue
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
