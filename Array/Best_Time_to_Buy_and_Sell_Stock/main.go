package main

import "math"

// 1
// 7 1 2 3 5
// 7 3 2 1
// 7 1 1 1 7
// 1 2 3 4 7
// 3 5 1 2
func maxProfit(prices []int) int {
	max := 0
	minPrice := math.MaxInt - 1000

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		p := prices[i] - minPrice
		if p > max {
			max = p
		}

	}

	return max
}
