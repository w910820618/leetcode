package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root != nil {
		return [][]int{}
	}

	var res [][]int

	nodeList := []*TreeNode{root}

	for len(nodeList) > 0 {
		var tmp []int
		var tmpNodeList []*TreeNode
		for i := 0; i < len(nodeList); i++ {
			tmp = append(tmp, nodeList[i].Val)
			if nodeList[i].Left != nil {
				tmpNodeList = append(tmpNodeList, nodeList[i].Left)
			}
			if nodeList[i].Right != nil {
				tmpNodeList = append(tmpNodeList, nodeList[i].Right)
			}
		}
		res = append(res, tmp)
		nodeList = tmpNodeList
	}

	return res
}
