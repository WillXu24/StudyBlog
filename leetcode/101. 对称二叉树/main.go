package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	res := isSymmetric(tree)
	fmt.Println(res)
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法思路：
// 如果根节点为空，返回true
// 如果根节点不为空，判断左右子树是否对称，调用递归函数
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

// 递归判断两个子树是否对称：
// 如果两个节点都为空，返回true
// 如果单个节点为空，返回false
// 如果两个节点相等，返回两个节点两个子树的判断结果的与
func helper(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return (a.Val == b.Val) && helper(a.Left, b.Right) && helper(a.Right, b.Left)
}
