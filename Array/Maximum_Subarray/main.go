package main

import "fmt"

func main() {
	r := maxSubArray([]int{-2, -1})
	fmt.Print(r)
}

//
// 1
// 1 2 3 4
// 6 -6 5 -6 -1 -5
// -5 1 2 -1 3 -5
// -5 1 2 -3 6 3 -5
// 6 -5 7 1
// -2 -1
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if max < sum {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return max
}
