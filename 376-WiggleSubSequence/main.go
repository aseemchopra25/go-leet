package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	prev := nums[1] - nums[0]
	count := 0
	if prev != 0 {
		count = 2
	} else {
		count = 1
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if (diff > 0 && prev <= 0) || (diff < 0 && prev >= 0) {
			count++
			prev = diff
		}
	}
	return count

}

func main() {
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	fmt.Println(wiggleMaxLength(nums))

}
