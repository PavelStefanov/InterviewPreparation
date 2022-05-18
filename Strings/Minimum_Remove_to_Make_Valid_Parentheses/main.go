package main

import "fmt"

func main() {
	r := minRemoveToMakeValid("))((")
	fmt.Print(r)
}

// ()
// ))((
// a(b())c)
//(abc
//abc)
// (a(b(c()))
// (a(b()c()))

type brace struct {
	char     byte
	position int
}

func minRemoveToMakeValid(s string) string {
	stack := []brace{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, brace{s[i], i})
		} else if s[i] == ')' {
			n := len(stack)
			if n > 0 {
				last := stack[n-1]
				if last.char == '(' {
					stack = stack[:n-1]
				} else {
					stack = append(stack, brace{s[i], i})
				}
			} else {
				stack = append(stack, brace{s[i], i})
			}
		}
	}

	res := []byte(s)

	for len(stack) > 0 {
		n := len(stack)
		last := stack[n-1]
		stack = stack[:n-1]

		res = remove(res, last.position)
	}

	return string(res)
}

func remove(slice []byte, s int) []byte {
	return append(slice[:s], slice[s+1:]...)
}
