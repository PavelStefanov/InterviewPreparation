package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3}
	r := uniqueOccurrences(arr)
	fmt.Print(r)
}

func uniqueOccurrences(arr []int) bool {
	sort.Sort(sort.IntSlice(arr))

	m := make(map[int]bool)

	cur, count := arr[0], 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == cur {
			count++
		} else {
			_, ok := m[count]
			if ok {
				return false
			} else {
				m[count] = true
			}

			count = 1
			cur = arr[i]
		}
	}

	_, ok2 := m[count]
	return !ok2
}
