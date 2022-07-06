package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	t := strconv.Itoa(x)
	l := len(t)
	for i := 0; i < l/2; i++ {
		if t[i] != t[l-1-i] {
			return false
		}
	}
	return true
}
func main() {

	fmt.Println(isPalindrome(10))
}
