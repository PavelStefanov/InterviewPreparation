package main

import "fmt"

func main() {
	r := maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	})
	fmt.Print(r)
}

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	visited := make(map[point]bool)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				size := calcIsland(grid, i, j, &visited)
				if size > max {
					max = size
				}
			}
		}
	}

	return max
}

type point struct {
	y int
	x int
}

func calcIsland(grid [][]int, y int, x int, visited *map[point]bool) int {
	p := point{y, x}

	_, ok := (*visited)[p]
	if ok {
		return 0
	}

	queue := []point{p}
	size := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		_, ok := (*visited)[cur]
		if grid[cur.y][cur.x] == 0 || ok {
			continue
		}

		size++
		(*visited)[cur] = true

		if cur.x > 0 {
			queue = append(queue, point{cur.y, cur.x - 1})
		}
		if cur.x < len(grid[cur.y])-1 {
			queue = append(queue, point{cur.y, cur.x + 1})
		}
		if cur.y > 0 {
			queue = append(queue, point{cur.y - 1, cur.x})
		}
		if cur.y < len(grid)-1 {
			queue = append(queue, point{cur.y + 1, cur.x})
		}

	}

	return size
}
