package main

import (
	"fmt"
)

func main() {
	r := shortestBridge([][]int{
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	})
	fmt.Print(r)
}

func shortestBridge(grid [][]int) int {
	n := len(grid)

	var y, x int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				y = i
				x = j
				break
			}
		}
	}

	dir := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	queueBridge := [][]int{}

	queueIsland1 := [][]int{{y, x}}
	grid[y][x] = 2

	for len(queueIsland1) > 0 {
		current := queueIsland1[0]
		queueIsland1 = queueIsland1[1:]

		queueBridge = append(queueBridge, []int{current[0], current[1]})

		for i := 0; i < len(dir); i++ {
			y = current[0] + dir[i][0]
			x = current[1] + dir[i][1]
			if y < n && x < n && y >= 0 && x >= 0 && grid[y][x] == 1 {
				queueIsland1 = append(queueIsland1, []int{y, x})
				grid[y][x] = 2
			}
		}
	}

	result := 0

	for len(queueBridge) > 0 {

		result++

		queueN := len(queueBridge)
		for i := 0; i < queueN; i++ {
			current := queueBridge[0]
			queueBridge = queueBridge[1:]

			for i := 0; i < len(dir); i++ {
				y = current[0] + dir[i][0]
				x = current[1] + dir[i][1]
				if y < n && x < n && y >= 0 && x >= 0 && grid[y][x] != 2 {

					if grid[y][x] == 1 {
						return result - 1
					}

					queueBridge = append(queueBridge, []int{y, x})
					grid[y][x] = 2
				}
			}
		}
	}

	return -1
}
