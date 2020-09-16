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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {

		return nil
	}

	var root int
	for i, v := range inorder {
		if v == preorder[0] {
			root = i
			break
		}
	}

	//preorderLeft, preorderRight := preorder[1:root+1], preorder[root+1:]
	//inorderLeft, inorderRight := inorder[:root], inorder[root+1:]

	//fmt.Println(preorderLeft, preorderRight)
	//fmt.Println(inorderLeft, inorderRight)

	return &TreeNode{
		Val: preorder[0],
		//Left:  buildTree(preorderLeft, inorderLeft),
		//Right: buildTree(preorderRight, inorderRight),
		Left:  buildTree(preorder[1:root+1], inorder[:root]),
		Right: buildTree(preorder[root+1:], inorder[root+1:]),
	}
}
