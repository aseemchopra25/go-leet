package main

import (
	"fmt"
)

func mySqrt(x int) int {
	l := 1
	r := x
	for l < r {
		m := (l + r + 1) / 2
		if m*m > x {
			r = m - 1
		} else {
			l = m
		}
	}
	return r

}

func main() {
	fmt.Println(mySqrt(2147483648))
}
