package main

import "fmt"

func main() {
	arr := []int{-1, 0, 5}
	r := search(arr, 2)
	fmt.Print(r)
}

func search(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s <= e {
		i := ((e - s) / 2) + s

		v := nums[i]
		if v == target {
			return i
		}

		if v > target {
			e = i - 1
		} else {
			s = i + 1
		}

	}

	return -1
}
