package main

func main() {

}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (helper(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func helper(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Val != b.Val {
		return false
	}
	return helper(a.Left, b.Left) && helper(a.Right, b.Right)
}
