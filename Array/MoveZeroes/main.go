package main

import "fmt"

// [1,2,3,4,5,6,7,8,9,10,11]
// 3
// [9,10,11,1,2,3,4,5,6,7,8]

func main() {
	arr := []int{1, 2, 0, 3, 0, 5, 6, 0, 0}
	moveZeroes(arr)
	fmt.Print(arr)
}

// 1,2,3,4,5
// 1,2,0,3,0,5,6,0,0
// 0,0,0,1,0,0,4,5
func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[l] != 0 && nums[r] != 0 {
			l++
			r++
		} else if nums[l] == 0 && nums[r] == 0 {
			r++
		} else {
			swap(nums, l, r)
			l++
		}
	}

}

func moveZeroes2(nums []int) {
	l, r := 0, 0
	for ; r < len(nums); r++ {
		if nums[r] != 0 {
			swap(nums, l, r)
			l++
		}
	}
}

func swap(nums []int, a int, b int) {
	t := nums[a]
	nums[a] = nums[b]
	nums[b] = t
}
