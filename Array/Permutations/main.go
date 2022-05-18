package main

import "fmt"

func main() {
	r := permute([]int{1, 2, 3})
	for _, v := range r {
		fmt.Print(v)
	}
}

func permute(nums []int) [][]int {
	result := [][]int{}

	recursive(nums, []int{}, &result)

	return result
}

func recursive(nums []int, current []int, result *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		current = append(current, nums[i])

		next := remove(nums, i)
		recursive(next, current, result)

		current = current[:len(current)-1]
	}
}

func remove(slice []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}
