package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, -1, 1, 1},
		{-1, -1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{-1, 1, 1, 1, 1},
	}

	re := cherryPickup(grid)

	fmt.Printf("%d", re)
}

type res struct {
	count   int
	canMove bool
}

type point struct {
	row int
	col int
}

func cherryPickup(grid [][]int) int {
	var points []point

	r := moveDown(grid, 0, 0, 0, points)
	if r.canMove == false {
		return 0
	}

	return r.count
}

func moveDown(grid [][]int, row int, col int, count int, points []point) (re res) {
	n := len(grid)

	localPoint := make([]point, len(points))
	copy(localPoint, points)

	re.count = count
	re.canMove = true

	if row >= n || col >= n || grid[row][col] == -1 {
		re.canMove = false
		return
	}

	if grid[row][col] == 1 && contains(localPoint, row, col) == false {
		re.count++
		localPoint = append(localPoint, point{row, col})
	}

	if row == n-1 && col == n-1 {
		pc := col - 1
		pr := row - 1

		u := moveUp(grid, pr, col, re.count, localPoint)
		l := moveUp(grid, row, pc, re.count, localPoint)

		if u.canMove == true && l.canMove == false {
			re.count = u.count
			return
		}

		if u.canMove == false && l.canMove == true {
			re.count = l.count
			return
		}

		if u.count > l.count {
			re.count = u.count
		} else {
			re.count = l.count
		}

		return
	}

	nc := col + 1
	nr := row + 1
	r := moveDown(grid, row, nc, re.count, localPoint)
	d := moveDown(grid, nr, col, re.count, localPoint)

	if r.canMove == false && d.canMove == false {
		re.canMove = false
		return
	}

	if r.canMove == true && d.canMove == false {
		re.count = r.count
		return
	}

	if r.canMove == false && d.canMove == true {
		re.count = d.count
		return
	}

	if r.count > d.count {
		re.count = r.count
	} else {
		re.count = d.count
	}

	return
}

func moveUp(grid [][]int, row int, col int, count int, points []point) (re res) {
	localPoint := make([]point, len(points))
	copy(localPoint, points)

	re.count = count
	re.canMove = true

	if row < 0 || col < 0 || grid[row][col] == -1 {
		re.canMove = false
		return
	}

	if grid[row][col] == 1 && contains(localPoint, row, col) == false {
		re.count++
		localPoint = append(localPoint, point{row, col})
	}

	if row == 0 && col == 0 {
		return
	}

	nc := col - 1
	nr := row - 1
	r := moveUp(grid, row, nc, re.count, localPoint)
	d := moveUp(grid, nr, col, re.count, localPoint)

	if r.canMove == false && d.canMove == false {
		re.canMove = false
		return
	}

	if r.canMove == true && d.canMove == false {
		re.count = r.count
		return
	}

	if r.canMove == false && d.canMove == true {
		re.count = d.count
		return
	}

	if r.count > d.count {
		re.count = r.count
	} else {
		re.count = d.count
	}

	return
}

func contains(points []point, row int, col int) bool {
	for _, v := range points {
		if v.row == row && v.col == col {
			return true
		}
	}
	return false
}
