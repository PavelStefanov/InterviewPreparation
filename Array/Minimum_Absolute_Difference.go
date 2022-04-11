package main

import (
	"fmt"
	"sort"
)

func main() {
	r := minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27})
	// r := minimumAbsDifference([]int{40, 11, 26, 27, -20})
	fmt.Print(r)
}

func minimumAbsDifference(arr []int) [][]int {
	min, res := -1, [][]int{}

	sort.Sort(sort.IntSlice(arr))

	for i := 0; i < len(arr)-1; i++ {
		d := arr[i+1] - arr[i]

		if min == -1 || d < min {
			min = d
			res = [][]int{{arr[i], arr[i+1]}}
		} else if min == d {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}

	return res
}
