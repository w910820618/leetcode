package main

import "fmt"

func main() {
	fmt.Println(numWays(7))
}

func numWays(n int) int {
	if n == 0 {
		return 1
	}

	if n <= 2 {
		return n
	}

	a := 1
	b := 2
	sum := 0

	for i := 2; i < n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}

	return sum
}
