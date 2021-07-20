package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
}

func missingNumber(nums []int) int {

	for i, v := range nums {
		if i != v {
			return i
		}
	}

	return len(nums)
}
