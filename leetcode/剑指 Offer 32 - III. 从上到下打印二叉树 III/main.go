package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val:   89,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := levelOrder(root)
	fmt.Println(res)
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// #树 #BFS 20201021
// 思路：在层次打印的基础上增加一个标志。
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	nodeQueue := []*TreeNode{root}
	sequent := false
	for len(nodeQueue) != 0 {
		levelLength := len(nodeQueue)
		temp := []int{}
		for i := 0; i < levelLength; i++ {
			node := nodeQueue[i]
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
			temp = append(temp, node.Val)
		}
		if sequent {

			for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}
			//fmt.Println(temp)
		}
		sequent = !sequent
		nodeQueue = nodeQueue[levelLength:]
		res = append(res, temp)
	}
	return res
}

// 思路：简单的BFS
// func levelOrder(root *TreeNode) []int {
// 	var res []int
// 	if root == nil {
// 		return res
// 	}

// 	nodeQueue := []*TreeNode{root}
// 	for len(nodeQueue) != 0 {
// 		node := nodeQueue[0]
// 		//fmt.Print(node.Val)
// 		nodeQueue = nodeQueue[1:]
// 		if node.Left != nil {
// 			nodeQueue = append(nodeQueue, node.Left)
// 		}
// 		if node.Right != nil {
// 			nodeQueue = append(nodeQueue, node.Right)
// 		}
// 		res = append(res, node.Val)
// 	}
// 	return res
// }
