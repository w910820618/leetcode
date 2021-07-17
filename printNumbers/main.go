package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(printNumbers(1))
}

func printNumbers(n int) []int {
	if n == 0 {
		return nil
	}

	max := math.Pow(10, float64(n))

	var result []int

	for i := 1; i < int(max); i++ {
		result = append(result, i)
	}

	return result
}
