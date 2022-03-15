package main

import "fmt"

func main() {
	arr := []int{1, 3}
	r := searchInsert(arr, 2)
	fmt.Print(r)
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	if target < nums[l] {
		return 0
	}

	if target > nums[r] {
		return r + 1
	}

	var p int
	for l <= r {
		p = l + ((r - l) / 2)

		if nums[p] == target {
			return p
		}

		if nums[p] > target {
			r = p - 1
		} else {
			l = p + 1
		}
	}

	if nums[p] > target {
		return p
	} else {
		return p + 1
	}
}
