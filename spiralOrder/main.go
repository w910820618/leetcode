package main

import "fmt"

func main() {
	//fmt.Println(spiralOrder([][]int{{3}, {2}}))
	fmt.Println(spiralOrder([][]int{{3}, {2}}))
	fmt.Println(spiralOrder([][]int{{3}, {2}}))
}

func spiralOrder(matrix [][]int) []int {
	var res []int

	if len(matrix) == 0 {
		return []int{}
	}

	left := 0
	right := len(matrix[0]) - 1
	low := 0
	high := len(matrix) - 1

	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[low][i])
		}

		low++

		if low > high {
			break
		}

		for i := low; i <= high; i++ {
			res = append(res, matrix[i][right])
		}

		right--

		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[high][i])
		}

		high--

		if high < low {
			break
		}

		for i := high; i >= low; i-- {
			res = append(res, matrix[i][left])
		}

		left++

		if left > right {
			break
		}
	}

	return res
}
