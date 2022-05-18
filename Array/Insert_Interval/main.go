package main

func main() {
	// insert([][]int{{1, 3}, {6, 9}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	// insert([][]int{{0, 5}, {9, 12}}, []int{7, 17})
	insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
}

// [[5,6]], [1,3]
// [[1,2]], [5,6]

// [[1,3]], [1,5]
// [[5,8]], [1,6]
// [[1,3], [7,8]] [5,6]
// [[1,3],[6,7]], [2,5]
// [[1,3],[6,7]], [3,5]
// [[1,2],[5,6],[7,8],[12,16]], [6,10]
func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}

	result := [][]int{}

	if newInterval[1] < intervals[0][0] {
		result = append([][]int{newInterval}, intervals...)
	}

	if newInterval[0] > intervals[n-1][1] {
		result = append(intervals, newInterval)
	}

	for i := 1; i < n; i++ {
		if newInterval[0] >= intervals[i-1][0] && newInterval[0] <= intervals[i][0] {
			result = append(result[:i+1], result[i:]...)
			result[i] = newInterval
		}
	}

	for i := 1; i < n; i++ {
		
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
