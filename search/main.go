package main

func main() {

}

func search(nums []int, target int) int {
	res := make(map[int]int)

	for _, v := range nums {
		res[v]++
	}

	return res[target]
}
