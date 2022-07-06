package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, x := range nums {
		if j, exists := m[target-x]; exists {
			return []int{j, i}
		}
		m[x] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
