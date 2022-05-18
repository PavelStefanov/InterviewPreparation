package main

import (
	"fmt"
)

func main() {
	r := shortestPathBinaryMatrix(
		[][]int{
			{0, 0, 0, 0},
			{1, 0, 0, 1},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		})
	fmt.Print(r)
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	n := len(grid)

	queue := [][]int{{0, 0, 1}}
	grid[0][0] = 1

	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current[0] == n-1 && current[1] == n-1 {
			return current[2]
		}

		for _, d := range dir {
			y := current[0] + d[0]
			x := current[1] + d[1]

			if y >= 0 && x >= 0 && y < n && x < n && grid[y][x] == 0 {
				queue = append(queue, []int{y, x, current[2] + 1})
				grid[y][x] = 1
			}
		}

	}

	return -1
}
