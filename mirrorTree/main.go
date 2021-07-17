package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	swap(root)

	return root
}

func swap(node *TreeNode) {
	if node == nil {
		return
	}

	tmp := node.Right
	node.Right = node.Left
	node.Left = tmp

	swap(node.Left)
	swap(node.Right)
}
