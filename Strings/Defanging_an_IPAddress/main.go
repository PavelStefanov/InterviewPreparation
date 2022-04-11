package main

import "fmt"

func main() {
	str := "255.100.50.0"
	r := defangIPaddr(str)
	fmt.Print(r)
}

func defangIPaddr(address string) string {
	chars := []rune(address)
	res := make([]rune, len(address)+6)

	i, j := 0, 0
	for i < len(chars) {
		if chars[i] == '.' {
			res[j] = '['
			j++
			res[j] = '.'
			j++
			res[j] = ']'
		} else {
			res[j] = chars[i]
		}

		i++
		j++
	}

	return string(res)
}
