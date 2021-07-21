package main

import "sort"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {

	nodeList := []*TreeNode{root}

	valList := make([]int, 0)

	for len(nodeList) != 0 {
		var tmpList []*TreeNode
		for i := 0; i < len(nodeList); i++ {
			valList = append(valList, nodeList[i].Val)
			if nodeList[i].Right != nil {
				tmpList = append(tmpList, nodeList[i].Right)
			}
			if nodeList[i].Left != nil {
				tmpList = append(tmpList, nodeList[i].Left)
			}
		}
		nodeList = tmpList
	}

	sort.Ints(valList)

	return valList[len(valList)-k]
}
