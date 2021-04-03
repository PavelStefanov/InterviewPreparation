package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 0, 0, 5},
		{4, 3, 1, 4},
		{0, 1, 1, 4},
		{1, 2, 1, 3},
		{0, 0, 1, 1},
	}

	setZeroes(matrix)

	print(matrix)
}

func setZeroes(matrix [][]int) {
	m := len(matrix)

	fc, fr := false, false

	for i := 0; i < m; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {

				if i == 0 {
					fr = true
				}
				if j == 0 {
					fc = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if fr == true {
		setRowZeros(matrix, 0)
	}
	if fc == true {
		setColumnZeros(matrix, 0)
	}
}

func setColumnZeros(matrix [][]int, column int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		matrix[i][column] = 0
	}
}

func setRowZeros(matrix [][]int, row int) {
	n := len(matrix[row])

	for i := 0; i < n; i++ {
		matrix[row][i] = 0
	}
}

func print(matrix [][]int) {
	for _, row := range matrix {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
	fmt.Println()
}
