package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	length := 0

	dumy := head

	for head != nil {
		length++
		head = head.Next
	}

	for i := 0; i < length-k; i++ {
		dumy = dumy.Next
	}

	return dumy
}
