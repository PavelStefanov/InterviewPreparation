package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(originalDigits("owoztneoer"))
}

func originalDigits(s string) string {
	var res string

	digit := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
		"4": "four",
		"5": "five",
		"6": "six",
		"7": "seven",
		"8": "eight",
		"9": "nine",
		"0": "zero",
	}

	for k := range digit {
		digit := []byte(digit[k])
		sort.Slice(digit, func(i int, j int) bool { return s[i] < s[j] })

		for {


			break
		}
	}



	return res
}
