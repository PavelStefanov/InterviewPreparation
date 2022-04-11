package main

import "fmt"

func main() {

	input := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(input)

	print(input)
}

func rotate(matrix [][]int) {
	n := len(matrix) - 1

	for i := 0; i < n; i++ {
		l := n - i
		if i == l {
			break
		}

		for j := 0; j < l-i; j++ {
			swap(matrix, i, i+j, l-j, i)
			swap(matrix, l-j, i, l, l-j)
			swap(matrix, l, l-j, i+j, l)
		}
	}
}

func swap(matrix [][]int, rowa int, columna int, rowb int, columnb int) {
	temp := matrix[rowa][columna]
	matrix[rowa][columna] = matrix[rowb][columnb]
	matrix[rowb][columnb] = temp
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
