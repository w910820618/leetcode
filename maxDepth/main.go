package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dep := 0

	nodeList := []*TreeNode{root}

	for len(nodeList) != 0 {
		var tmpList []*TreeNode
		for i := 0; i < len(nodeList); i++ {
			if nodeList[i].Left != nil {
				tmpList = append(tmpList, nodeList[i].Left)
			}
			if nodeList[i].Right != nil {
				tmpList = append(tmpList, nodeList[i].Right)
			}
		}
		dep++
		nodeList = tmpList
	}

	return dep
}
