package main

import "fmt"

func main() {
	image := [][]int{
		[]int{1, 1, 1},
		[]int{1, 1, 0},
		[]int{1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2

	r := floodFill(image, sr, sc, newColor)
	fmt.Print(r)
}

type point struct {
	y int
	x int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]

	queue := []point{}
	m := make(map[point]bool)

	queue = append(queue, point{sr, sc})

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		_, ok := m[cur]
		if ok {
			continue
		}

		m[cur] = true
		image[cur.y][cur.x] = newColor

		// y
		if cur.y+1 < len(image) && image[cur.y+1][cur.x] == color {
			queue = append(queue, point{cur.y + 1, cur.x})
		}
		if cur.y-1 >= 0 && image[cur.y-1][cur.x] == color {
			queue = append(queue, point{cur.y - 1, cur.x})
		}

		// x
		if cur.x+1 < len(image[cur.y]) && image[cur.y][cur.x+1] == color {
			queue = append(queue, point{cur.y, cur.x+1})
		}
		if cur.x-1 >= 0 && image[cur.y][cur.x-1] == color {
			queue = append(queue, point{cur.y, cur.x-1})
		}
	}

	return image
}
