package main

import (
	"fmt"
)

func main() {
	r := minCostToMoveChips([]int{2, 2, 2, 3, 3})
	fmt.Print(r)
}

func minCostToMoveChips(position []int) int {
	even := 0
	odd := 0

	for i := 0; i < len(position); i++ {
		if position[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even > odd {
		return odd
	} else {
		return even
	}
}
