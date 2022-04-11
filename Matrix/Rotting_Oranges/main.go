package main

import "fmt"

func main() {
	r := orangesRotting([][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1}})
	fmt.Print(r)
}

func orangesRotting(grid [][]int) int {
	// only fresh oranges. Return -1
	// only rotten return 0
	// find the longest way from rotten to fresh

	maxWay := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				m := make(map[point]bool)
				queue := []point{}

				

				hasFresh = true
			} else if grid[i][j] == 2 {
				hasRotten = true
			}
		}
	}

	return maxWay

}

type point struct {
	y int
	x int
}
