package main

import "fmt"

func main() {
	a := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	res := HeapSort(a)
	fmt.Println(res)
}

// HeapSort 插入排序
func HeapSort(a []int) []int {
	// 遍历所有非叶子节点，从最接近叶子节点的非叶子节点遍历到根节点
	for i := len(a) / 2; i >= 0; i-- {
		// 堆排序
		heapAdjust(a, i)
	}
	// 从倒数第二个节点开始循环遍历
	for i := len(a) - 1; i > 0; i-- {
		// 将最大值调换到序列尾部
		a[0], a[i] = a[i], a[0]
		// 堆排序
		heapAdjust(a[:i], 0)
	}
	return a
}

// 堆排序：将最大数调换到根节点
func heapAdjust(a []int, root int) {
	// 顺着节点找到最小的位置（不太好表达
	for i := 2*root + 1; i < len(a); i = i*2 + 1 {
		// 如果右子树存在，找出左右子树最大值下标
		if i+1 < len(a) && a[i] < a[i+1] {
			i++
		}
		// 如果根节点大于子树，跳出循环
		if a[root] > a[i] {
			break
		}
		// 根节点调换为最值
		a[root], a[i] = a[i], a[root]
		// 原本的根节点变为i，进行下一次比较，直到调换到最小的位置
		root = i
	}
}
