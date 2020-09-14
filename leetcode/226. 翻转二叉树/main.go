package main

import "fmt"

// TreeNode 二叉树
type TreeNode struct {
	    Val int
        Left *TreeNode
	    Right *TreeNode
}

func main()  {
	tree:=TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:2,
			Left: &TreeNode{
				Val:1,
				
			},
			Right: &TreeNode{
				Val:3,
			},
		},
		Right: &TreeNode{
			Val:7,
			Left: &TreeNode{
				Val:6,
				
			},
			Right: &TreeNode{
				Val:9,
			},
		},
	}

	res:=invertTree(&tree)

	fmt.Println("     ",res.Val)
	fmt.Println(" ",res.Left.Val,"      ",res.Right.Val)
	fmt.Println(res.Left.Left.Val," ",res.Left.Right.Val,"  ",res.Right.Left.Val," ",res.Right.Right.Val)
}

func invertTree(root *TreeNode) *TreeNode {
    // 递归结束条件
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	// 左右节点分别递归
	invertTree(root.Right)
	invertTree(root.Left)
	return root
}