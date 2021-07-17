package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1. 如何记录前一个节点
2. 如何记录当前节点
*/

func reverseList(head *ListNode) *ListNode {

	cur := head
	var pre *ListNode

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}
