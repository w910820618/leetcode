package main

import "fmt"

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		if nums[left]%2 == 0 && nums[right]%2 != 0 {
			center := nums[right]
			nums[right] = nums[left]
			nums[left] = center
			right--
			left++
		} else if nums[left]%2 != 0 {
			left++
		} else if nums[right]%2 == 0 {
			right--
		}
	}

	return nums
}
