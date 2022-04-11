package main

import "fmt"

// [1,2,3,4,5,6,7,8,9,10,11]
// 3
// [9,10,11,1,2,3,4,5,6,7,8]
//

// [9,10,11,1,2,3,7,8,1,2,3]
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(arr, 3)
	fmt.Print(arr)
}

func rotate1(nums []int, k int) {
	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}
	rotate := make([]int, k)
	arr := make([]int, l-k)

	copy(rotate[:], nums[l-k:])
	copy(arr[:], nums[:l-k])

	for i := 0; i < len(rotate); i++ {
		nums[i] = rotate[i]
	}
	for i := 0; i < len(arr); i++ {
		nums[i+k] = arr[i]
	}
}

func rotate2(nums []int, k int) {
	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}

	for i := 0; i < k; i++ {
		last := nums[l-1]

		for j := l - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}

		nums[0] = last
	}
}

func swap(nums []int, a int, b int) {
	c := nums[a]
	nums[a] = nums[b]
	nums[b] = c
}
