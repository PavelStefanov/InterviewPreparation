package main

import "fmt"

func main() {
	r := combine(4, 2)
	for _, v := range r {
		fmt.Print(v)
	}
}

func combine(n int, k int) [][]int {
	res := [][]int{}

	backtrack(1, n, k, []int{}, &res)

	return res
}

func backtrack(start, n, k int, current []int, result *[][]int) {
	if len(current) == k {
		n := make([]int, len(current))
		copy(n, current)
		*result = append(*result, n)
		return
	}

	for i := start; i <= n; i++ {
		current = append(current, i)
		backtrack(i+1, n, k, current, result)
		current = current[:len(current)-1]
	}
}
