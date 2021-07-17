package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		head = l2
		l2 = l2.Next
	} else {
		head = l1
		l1 = l1.Next
	}

	pre := head

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			pre.Next = l2
			l2 = l2.Next
		} else {
			pre.Next = l1
			l1 = l1.Next
		}
		pre = pre.Next
	}

	if l1 != nil {
		pre.Next = l1
	} else if l2 != nil {
		pre.Next = l2
	}

	return head
}
