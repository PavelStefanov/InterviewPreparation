package main

import "fmt"

func main() {
	r := numRookCaptures([][]byte{

		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	})
	fmt.Print(r)
}

func numRookCaptures(board [][]byte) int {
	r := []int{}
	c := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				r = []int{i, j}
			}
		}
	}

	// right
	for j := r[1]; j < len(board[r[0]]); j++ {
		if board[r[0]][j] == 'B' {
			break
		}
		if board[r[0]][j] == 'p' {
			c++
			break
		}
	}

	// left
	for j := r[1]; j >= 0; j-- {
		if board[r[0]][j] == 'B' {
			break
		}
		if board[r[0]][j] == 'p' {
			c++
			break
		}
	}

	// down
	for i := r[0]; i < len(board); i++ {
		if board[i][r[1]] == 'B' {
			break
		}
		if board[i][r[1]] == 'p' {
			c++
			break
		}
	}

	// top
	for i := r[0]; i >= 0; i-- {
		if board[i][r[1]] == 'B' {
			break
		}
		if board[i][r[1]] == 'p' {
			c++
			break
		}
	}

	return c
}
