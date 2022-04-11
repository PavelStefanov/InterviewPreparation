package main

import (
	"fmt"
	"math"
)

func main() {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	r := updateMatrix(mat)
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
	}
}

func updateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	for i := 0; i < len(res); i++ {
		res[i] = newInts(len(mat[i]), math.MaxInt)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
			} else {
				for g := 0; g < len(mat); g++ {
					for h := 0; h < len(mat[g]); h++ {
						if mat[g][h] == 0 {
							res[i][j] = min(res[i][j], int(math.Abs(float64(i-g))+math.Abs(float64(j-h))))
						}
					}
				}
			}
		}
	}

	return res
}

func updateMatrix2(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	for i := 0; i < len(res); i++ {
		res[i] = newInts(len(mat[i]), math.MaxInt-100000)
	}

	// from top
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
			} else {
				if i > 0 {
					res[i][j] = min(res[i][j], res[i-1][j]+1)
				}
				if j > 0 {
					res[i][j] = min(res[i][j], res[i][j-1]+1)
				}
			}
		}
	}

	//from bot
	for i := len(res) - 1; i >= 0; i-- {
		for j := len(res[i]) - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				res[i][j] = 0
			} else {
				if i < len(res)-1 {
					res[i][j] = min(res[i][j], res[i+1][j]+1)
				}
				if j < len(res[i])-1 {
					res[i][j] = min(res[i][j], res[i][j+1]+1)
				}
			}
		}
	}

	return res
}

func newInts(n, v int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = v
	}
	return s
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
