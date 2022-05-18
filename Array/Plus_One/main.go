package main

func plusOne(digits []int) []int {
	add := 1

	for i := len(digits) - 1; i >= 0; i-- {
		s := digits[i] + add
		add = 0

		if s <= 9 {
			digits[i] = s
		} else {
			digits[i] = 0
			add = 1
		}

		if add == 0 {
			break
		}
	}

	if add == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
