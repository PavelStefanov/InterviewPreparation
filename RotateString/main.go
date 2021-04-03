package main

import "fmt"

func main() {

	a := "ffu"
	b := "fuf"

	res := rotateString(a, b)

	fmt.Printf("%t", res)

}

func rotateString(A string, B string) bool {
	if len(B) != len(A) {
		return false
	}

	if len(B) == 0 && len(A) == 0 {
		return true
	}

	str := A + A

	return isSubsequence(str, B)
}

func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == t[j] {
			j++
		} else {
			j = 0
			if s[i] == t[j] {
				j++
			}
		}

		if j == len(t) {
			return true
		}
	}

	return false
}
