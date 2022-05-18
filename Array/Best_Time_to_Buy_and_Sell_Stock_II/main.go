package main

import "fmt"

func main() {
	r := maxProfit([]int{6, 1, 3, 2, 4, 7})
	fmt.Print(r)
}

// 1
// 6 5 4 3 2 1
// 1 2 3 4 5 6
// 7 1 5 2 5 6
// 7 1 2 3 4
// 1 2 2 3 3 4 4 4 5 5 5
// 1 3 2 4
// 6,1,3,2,4,7
func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
