package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return swap(root.Left, root.Right)
}

func swap(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if (left != nil && right == nil) || (left == nil && right != nil) || (left.Val != right.Val) {
		return false
	}

	return swap(left.Left, right.Right) && swap(left.Right, right.Left)
}
