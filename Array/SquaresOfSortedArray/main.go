package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 3}
	r := sortedSquares(arr)
	fmt.Print(r)
}

func sortedSquares(nums []int) []int {
	len := len(nums)

	l, r, i := 0, len-1, len-1

	res := make([]int, len)
	for l <= r {
		if math.Abs(float64(nums[l])) > math.Abs(float64(nums[r])) {
			res[i] = int(math.Pow(float64(nums[l]), 2))
			l++
		} else {
			res[i] = int(math.Pow(float64(nums[r]), 2))
			r--
		}
		i--
	}

	return res
}
