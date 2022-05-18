package main

import "fmt"

func main() {
	r := rob2([]int{2, 1, 1, 2, 1, 1, 3, 9, 3})
	fmt.Print(r)
}

// 2 1 1 2
// 4 5 3 1
// 4 5 3 1 1
// 1 3 3 1 5
// 1 3 3 4 1 1 1
// 1 3 3 4 9 1 1
// 2 1 1 2 9 1
// 2 1 1 2 1 1 3 9 3
// 1
// 1 2
// 2 1
// 2 1 2
// 1 3 1

func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	if n == 3 {
		return max(nums[0]+nums[2], nums[1])
	}

	sum1 := nums[0]

	// [2,7,9,3,1]
	fmt.Print("sum1")
	fmt.Println()
	i := 0
	for i < n-3 {
		b := nums[i+3]

		c := nums[i+2]
		if n-1 >= i+4 {
			c += nums[i+4]
		}

		if b >= c {
			sum1 += nums[i+3]
			i = i + 3
			fmt.Print(i)
		} else {
			sum1 += nums[i+2]
			i = i + 2
			fmt.Print(i)
		}
	}

	sum2 := nums[1]

	fmt.Println()
	fmt.Print("sum2")
	fmt.Println()
	i = 1
	for i < n-3 {
		b := nums[i+3]

		c := nums[i+2]
		if n-1 >= i+4 {
			c += nums[i+4]
		}

		if b > c {
			sum2 += nums[i+3]
			i = i + 3
			fmt.Print(i)
		} else {
			sum2 += nums[i+2]
			i = i + 2
			fmt.Print(i)
		}
	}

	return max(sum1, sum2)
}

func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	d := make([]int, n)
	d[0] = nums[0]
	d[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		d[i] = max(d[i-1], d[i-2]+nums[i])
	}

	return d[n-1]
}

func robRec(nums []int, position int, memory *map[int]int) int {
	n := len(nums)

	if position > n-1 {
		return 0
	}
	if position >= n-2 {
		return nums[position]
	}

	v, ok := (*memory)[position]
	if ok {
		return v
	}

	a := robRec(nums, position+2, memory)
	b := robRec(nums, position+3, memory)

	return nums[position] + max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
