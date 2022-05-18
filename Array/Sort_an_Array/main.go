package main

import "fmt"

func main() {
	r := sortArray([]int{5, 1, 1, 2, 0, 0})
	fmt.Print(r)
}

func sortArray(nums []int) []int {
	// return quickSort(nums)
	return mergeSort(nums)
}

func quickSort(nums []int) []int {
	quickSortNested(&nums, 0, len(nums)-1)
	return nums
}

func quickSortNested(nums *[]int, s int, e int) {
	if s >= e {
		return
	}

	pivot := (*nums)[e]
	min := s

	for i := s; i <= e-1; i++ {
		if (*nums)[i] <= pivot {
			swap(nums, min, i)
			min++
		}
	}

	swap(nums, min, e)

	quickSortNested(nums, s, min-1)
	quickSortNested(nums, min+1, e)
}

func mergeSort(nums []int) []int {
	return mergeSortNested(nums, 0, len(nums)-1)
}

func mergeSortNested(nums []int, s int, e int) []int {
	if s == e {
		return []int{nums[s]}
	}

	m := (e-s)/2 + s

	arr1 := mergeSortNested(nums, s, m)
	arr2 := mergeSortNested(nums, m+1, e)
	arr1P := 0
	arr2P := 0

	newArr := make([]int, len(arr1)+len(arr2))

	for i := 0; i < len(newArr); i++ {
		if arr1P < len(arr1) && arr2P < len(arr2) {
			if arr1[arr1P] < arr2[arr2P] {
				newArr[i] = arr1[arr1P]
				arr1P++
			} else {
				newArr[i] = arr2[arr2P]
				arr2P++
			}
		} else {
			if arr1P < len(arr1) {
				newArr[i] = arr1[arr1P]
				arr1P++
			} else {
				newArr[i] = arr2[arr2P]
				arr2P++
			}
		}
	}

	return newArr
}

func swap(nums *[]int, a int, b int) {
	t := (*nums)[a]
	(*nums)[a] = (*nums)[b]
	(*nums)[b] = t
}
