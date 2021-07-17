package main

func main() {

}

func majorityElement(nums []int) int {
	length := (len(nums) - 1) / 2

	count := make(map[int]int)

	for _, v := range nums {
		count[v]++
		if count[v] >= length {
			return v
		}
	}

	return -1
}
