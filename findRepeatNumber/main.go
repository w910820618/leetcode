package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

func findRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return -1
}
