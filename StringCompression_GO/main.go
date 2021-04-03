package main

import (
	"fmt"
)

func main() {

	inputs := [][]byte{
		{'a'},
		{'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b'},
		{'a', 'b', 'c', 'd'},
		{'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c'},
	}

	for _, input := range inputs {
		fmt.Println(string(input))
		res := compress(input)
		fmt.Println(string(input))
		fmt.Println(res)
	}
}

func compress(chars []byte) int {
	if len(chars) < 2 {
		return len(chars)
	}

	l := 0
	var lc byte

	for i := 0; i < len(chars); {
		lc = chars[i]
		count := 0
		for i < len(chars) && lc == chars[i] {
			count++
			i++
		}

		if chars[l] != lc {
			chars[l] = lc
		}

		if count == 1 {
			l++
		} else {
			s := fmt.Sprint(count)
			for c := 0; c < len(s); c++ {
				chars[l+1+c] = s[c]
			}

			l = l + len(s) + 1
		}
	}
	return l
}
