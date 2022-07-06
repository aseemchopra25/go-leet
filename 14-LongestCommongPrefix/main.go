package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	s := ""
	minlen := 201
	for i := range strs {
		if len(strs[i]) < minlen {
			minlen = len(strs[i])
		}
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for j := 0; j < minlen; j++ {
		flag := 1
		temp := ""
		for i := 1; i < len(strs); i++ {
			if string(strs[i][j]) != string(strs[i-1][j]) {
				flag = 0
				return s
			}
			temp = string(strs[i][j])
		}
		if flag == 1 {
			s += temp
		}
	}

	return s
}

func main() {
	strs := []string{"car", "cir"}
	fmt.Println(longestCommonPrefix(strs))
}
