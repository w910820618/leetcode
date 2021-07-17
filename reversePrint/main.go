package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var list []int

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	var res []int
	for i := len(list) - 1; i >= 0; i-- {
		res = append(res, list[i])
	}

	return res
}
