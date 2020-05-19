package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSym(root, root)
}

func isSym(r1, r2 *TreeNode) bool {
	if r1 == nil || r2 == nil {
		return r1 == r2
	}
	if r1.Val != r2.Val {
		return false
	}
	return isSym(r1.Left, r2.Right) && isSym(r1.Right, r2.Left)
}
