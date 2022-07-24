package main

import "fmt"

func main() {
	r := addBinary("11", "1")
	fmt.Println(r)
}

func addBinary(a string, b string) string {
	r := ""

	i := len(a) - 1
	j := len(b) - 1

	rest := 0
	for i >= 0 && j >= 0 {
		if i >= 0 && j >= 0 {
			if a[i] == '1' && b[j] == '1' {
				if rest == 1 {
					r = "1" + r
				} else {
					r = "0" + r
					rest = 1
				}
			} else if a[i] == '0' && b[j] == '0' {
				if rest == 1 {
					r = "1" + r
					rest = 0
				} else {
					r = "0" + r
				}
			} else {
				if rest == 1 {
					r = "0" + r
				} else {
					r = "1" + r
				}
			}
		}
		i--
		j--
	}

	for i >= 0 {
		if a[i] == '1' {
			if rest == 1 {
				r = "0" + r
			} else {
				r = "1" + r
			}
		} else {
			if rest == 1 {
				r = "1" + r
				rest = 0
			} else {
				r = "0" + r
			}
		}
		i--
	}

	for j >= 0 {
		if b[j] == '1' {
			if rest == 1 {
				r = "0" + r
			} else {
				r = "1" + r
			}
		} else {
			if rest == 1 {
				r = "1" + r
				rest = 0
			} else {
				r = "0" + r
			}
		}
		j--
	}

	if rest == 1 {
		r = "1" + r
	}

	return r
}
