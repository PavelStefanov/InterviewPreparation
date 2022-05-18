package main

import "fmt"

func main() {
	r := minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
	fmt.Print(r)
}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minimumTotalRecursively(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	m := make(map[point]int)

	left := getMinimumTotal(triangle, 1, 0, &m)
	right := getMinimumTotal(triangle, 1, 1, &m)

	if left > right {
		return triangle[0][0] + right
	} else {
		return triangle[0][0] + left
	}
}

func getMinimumTotal(triangle [][]int, row int, position int, m *map[point]int) int {
	if row == len(triangle)-1 {
		return triangle[row][position]
	}

	v, ok := (*m)[point{row, position}]
	if ok {
		return v
	}

	left := getMinimumTotal(triangle, row+1, position, m)
	right := getMinimumTotal(triangle, row+1, position+1, m)

	if left > right {
		(*m)[point{row, position}] = triangle[row][position] + right
		return triangle[row][position] + right
	} else {
		(*m)[point{row, position}] = triangle[row][position] + left
		return triangle[row][position] + left
	}
}

type point struct {
	row int
	col int
}
