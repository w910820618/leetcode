package main

import "fmt"

func main() {
	fmt.Println(byte(firstUniqChar("abaccdeff")))
}

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}

	tmp := make(map[byte]int)

	for _, v := range s {
		tmp[byte(v)]++
	}

	for _, v := range s {
		if tmp[byte(v)] == 1 {
			return byte(v)
		}
	}

	return ' '
}
