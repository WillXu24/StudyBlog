/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import "fmt"

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	res := buildTree(preorder, inorder)

	// 打印前两层（待改善
	fmt.Println(res.Val)
	fmt.Println(res.Left.Val, res.Right.Val)
}

// 解法
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 设定返回条件
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 分离左右子树
	var root int
	for i, v := range inorder {
		if v == preorder[0] {
			root = i
			break
		}
	}

	// 左右子树
	// pre_left,pre_right:=preorder[1:root+1],preorder[root+1:]
	// in_left,in_right:=inorder[0:root],inorder[root+1:]

	// 递归
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:root+1], inorder[0:root]),
		Right: buildTree(preorder[root+1:], inorder[root+1:]),
	}
}
