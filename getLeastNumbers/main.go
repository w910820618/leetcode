package main

import "sort"

func main() {

}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	sort.Ints(arr)

	return arr[0:k]
}
