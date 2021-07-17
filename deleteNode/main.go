package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	dumy := head

	for dumy.Next != nil {
		cur := dumy.Next
		if cur.Val == val {
			dumy.Next = cur.Next
		}
		dumy = cur
	}

	return head
}
