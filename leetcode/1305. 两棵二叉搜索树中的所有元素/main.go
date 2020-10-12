package main

import "fmt"

func main() {
	r1 := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	r2 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	res := getAllElements(&r1, &r2)
	fmt.Println(res)
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var a, b []int
	rangeTree(root1, &a)
	rangeTree(root2, &b)
	//fmt.Println(a,b)
	res := mergeSort(a, b)
	return res
}

// 中序遍历返回有序数组
func rangeTree(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}
	rangeTree(root.Left, a)
	*a = append(*a, root.Val)
	rangeTree(root.Right, a)
}

// 有序数组的归并排序
func mergeSort(a, b []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}
