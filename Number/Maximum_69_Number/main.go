package main

import (
	"fmt"
	"math"
)

func main() {
	r := maximum69Number(9669)
	fmt.Print(r)
}

func maximum69Number(num int) int {
	t := num

	c, i := 1, 0
	for t > 0 {
		n := t % 10

		if n == 6 {
			i = c
		}
		t = t / 10
		c++
	}

	if i == 1 {
		num = num + 3
	} else if i > 1 {
		d := int(math.Pow(float64(10), float64(i-1)))
		num = num + (d * 3)
	}

	return num
}
