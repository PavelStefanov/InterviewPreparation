package main

import "fmt"

// 1. 1
// 2 2. 4
// 2 3 2. 4
// 1 2 3. 4
// 3 2 1. 4
// 1 2 3. 7
func main() {
	r := twoSum([]int{}, 9)
	fmt.Print(r)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		v, ok := m[target-nums[i]]
		if ok {
			return []int{i, v}
		} else {
			m[nums[i]] = i
		}
	}

	return []int{}
}
