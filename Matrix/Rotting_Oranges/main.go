package main

import (
	"fmt"
	"math"
)

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

	freshes := make(map[point]value)
	rottenes := []point{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				freshes[point{i, j}] = value{false, math.MaxInt - 1000}
			} else if grid[i][j] == 2 {
				rottenes = append(rottenes, point{i, j})
			}
		}
	}

	if len(freshes) == 0 {
		return 0
	}

	for _, r := range rottenes {
		bfs(&grid, &freshes, r.y, r.x)
	}

	maxPath := -1

	for _, v := range freshes {
		if !v.visited {
			return -1
		} else if v.path > maxPath {
			maxPath = v.path
		}
	}

	return maxPath
}

func bfs(grid *[][]int, freshes *map[point]value, i int, j int) {
	n, m := len(*grid), len((*grid)[0])

	queue := []path{{i, j, 0}}
	visited := make(map[point]bool)
	visited[point{i, j}] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		v, ok := (*freshes)[point{cur.y, cur.x}]
		if ok {
			path := v.path
			if path > cur.dest {
				path = cur.dest
			}
			(*freshes)[point{cur.y, cur.x}] = value{true, path}
		}

		// up
		if cur.y > 0 && (*grid)[cur.y-1][cur.x] != 0 {
			_, ok := visited[point{cur.y - 1, cur.x}]
			if !ok {
				queue = append(queue, path{cur.y - 1, cur.x, cur.dest + 1})
				visited[point{cur.y - 1, cur.x}] = true
			}
		}

		// down
		if cur.y < n-1 && (*grid)[cur.y+1][cur.x] != 0 {
			_, ok := visited[point{cur.y + 1, cur.x}]
			if !ok {
				queue = append(queue, path{cur.y + 1, cur.x, cur.dest + 1})
				visited[point{cur.y + 1, cur.x}] = true
			}
		}

		// left
		if cur.x > 0 && (*grid)[cur.y][cur.x-1] != 0 {
			_, ok := visited[point{cur.y, cur.x - 1}]
			if !ok {
				queue = append(queue, path{cur.y, cur.x - 1, cur.dest + 1})
				visited[point{cur.y, cur.x - 1}] = true
			}
		}

		// right
		if cur.x < m-1 && (*grid)[cur.y][cur.x+1] != 0 {
			_, ok := visited[point{cur.y, cur.x + 1}]
			if !ok {
				queue = append(queue, path{cur.y, cur.x + 1, cur.dest + 1})
				visited[point{cur.y, cur.x + 1}] = true
			}
		}
	}
}

type point struct {
	y int
	x int
}

type path struct {
	y    int
	x    int
	dest int
}

type value struct {
	visited bool
	path    int
}
