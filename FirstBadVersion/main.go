package main

import "fmt"

func main() {
	r := firstBadVersion(5)
	fmt.Print(r)
}

func firstBadVersion(n int) int {
	l, r := 1, n
	b := 1
	for l <= r {
		p := l + ((r - l) / 2)

		isBad := isBadVersion(p)
		if isBad {
			b = p
		}

		if isBad == true {
			r = p - 1
		} else {
			l = p + 1
		}
	}

	return b
}

func isBadVersion(version int) bool {
	if version == 4 {
		return true
	} else {
		return false
	}
}
