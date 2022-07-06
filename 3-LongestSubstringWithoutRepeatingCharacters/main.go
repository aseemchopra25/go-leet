package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 1
	l := len(s)
	for i := 0; i < l; i++ {
		m := make(map[string]int)
		m[string(s[i])] = 1
		temp := 1
		for j := i + 1; j < l; j++ {
			if _, ok := m[string(s[j])]; ok {
				break
			} else {
				m[string(s[j])] = 1
				temp++
			}
		}
		if temp > res {
			res = temp
		}
	}
	return res

}

func main() {
	fmt.Println(lengthOfLongestSubstring("asdaeesasdafewfwfdsdcds"))
}
