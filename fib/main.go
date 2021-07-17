package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}

	a := 1
	b := 0
	sum := 0

	for i := 1; i < n; i++ {
		sum = (a + b) % 1000000007
		b = sum
		a = b
	}

	return sum
}

func main() {
	fmt.Println(fib(2))
}
