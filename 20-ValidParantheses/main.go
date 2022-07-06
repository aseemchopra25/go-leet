package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	for len(s) != 0 {
		if strings.Contains(s, "()") {
			s = strings.Replace(s, "()", "", -1)
		} else if strings.Contains(s, "[]") {
			s = strings.Replace(s, "[]", "", -1)

		} else if strings.Contains(s, "{}") {
			s = strings.Replace(s, "{}", "", -1)

		} else {
			return false
		}
	}
	if len(s) == 0 {
		return true
	}
	return true
}

func main() {

	fmt.Println(isValid("(){}()]"))

}
