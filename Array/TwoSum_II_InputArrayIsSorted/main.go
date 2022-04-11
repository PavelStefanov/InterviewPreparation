package main

import "fmt"

// [1,2,3,4,5,6,7,8,9,10,11]
// 3
// [9,10,11,1,2,3,4,5,6,7,8]

func main() {
	arr := []int{2, 7, 11, 15}
	r := twoSum2(arr, 9)
	fmt.Print(r)
}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		add := target - numbers[i]
		ser := binarySearch(numbers, i+1, len(numbers)-1, add)
		if ser != -1 {
			return []int{i + 1, ser + 1}
		}
	}

	return nil
}

func twoSum2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		s := numbers[l] + numbers[r]
		if s == target {
			return []int{l + 1, r + 1}
		} else if s < target {
			l++
		} else {
			r++
		}
	}
	return nil
}

func binarySearch(nums []int, l int, r int, target int) int {
	res := -1

	for l <= r {
		p := ((r - l) / 2) + l

		if nums[p] == target {
			res = p
			break
		}

		if nums[p] < target {
			l = p + 1
		} else {
			r = p - 1
		}
	}

	return res
}
