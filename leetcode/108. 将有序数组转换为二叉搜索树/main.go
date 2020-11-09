package main

import "fmt"

func main() {
	a := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(a)
	level := levelOrder(res)
	fmt.Println(level)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{Val: nums[mid], Left: sortedArrayToBST(nums[:mid]), Right: sortedArrayToBST(nums[mid+1:])}
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	//var levelCount int
	nodeQueue := []*TreeNode{root}
	for len(nodeQueue) != 0 {
		levelLength := len(nodeQueue)
		temp := []int{}
		for i := 0; i < levelLength; i++ {
			node := nodeQueue[i]
			//fmt.Print(node.Val)
			//nodeQueue = nodeQueue[1:]
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
			temp = append(temp, node.Val)
		}
		nodeQueue = nodeQueue[levelLength:]
		res = append(res, temp)
	}
	return res
}
