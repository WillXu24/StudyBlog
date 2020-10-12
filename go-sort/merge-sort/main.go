package main

import "fmt"

func main() {
	a := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	res := MergeSort(a)
	fmt.Println(res)
}

// MergeSort 插入排序
func MergeSort(a []int) []int {
	// 递归退出条件
	if len(a) < 2 {
		return a
	}
	// 二分序列
	m := len(a) / 2
	// 左右半序列分别递归排序
	l := MergeSort(a[:m])
	r := MergeSort(a[m:])
	// 将左右序列归并
	return merge(l, r)
}

func merge(a, b []int) []int {
	var res []int
	i, j := 0, 0
	// 双指针序列对比拷贝到新序列
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	// 将最后剩余部分拷贝到新序列结尾，这里有一个序列其实已经到了结尾，实际上只拷贝了一个序列
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}
