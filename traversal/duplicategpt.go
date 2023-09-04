package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func main() {
	nums := []int{1, 1, 2, 2, 2, 2, 2, 3, 4, 4, 5, 5}
	removeDuplicates(nums)
	fmt.Println(nums)
}
