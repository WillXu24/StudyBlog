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

// 迭代思路：
// 如果根节点为空，返回true
// 如果根节点不为空，返回调用函数结果
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

// 迭代函数：
// 采用队列，
func helper(a, b *TreeNode) bool {
	nodeQueue := []*TreeNode{a, b}
	for len(nodeQueue) > 0 {
		node1, node2 := nodeQueue[0], nodeQueue[1]
		nodeQueue = nodeQueue[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		nodeQueue = append(nodeQueue, node1.Left, node2.Right)
		nodeQueue = append(nodeQueue, node1.Right, node2.Left)
	}
	return true
}
