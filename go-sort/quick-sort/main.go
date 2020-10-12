package main

import "fmt"

func main() {
	a := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	res := QuickSort(a)
	fmt.Println(res)
}

// QuickSort 插入排序
func QuickSort(a []int) []int {
	qSort(a, 0, len(a)-1)
	return a
}

// 对a序列快速排序
func qSort(a []int, left, right int) {
	// 递归退出条件
	if left >= right {
		return
	}
	// 寻找并确定枢轴位置
	p := partition(a, left, right)
	// 对枢轴两边的序列分别进行快速排序
	qSort(a, left, p-1)
	qSort(a, p+1, right)
}

// 寻找并确定枢轴位置
func partition(a []int, left, right int) int {
	// 选择首位为枢轴
	p := a[left]
	// 双指针逼近，重合时退出
	for left != right {
		// 最右值大于枢轴，右指针左移
		for left < right && p < a[right] {
			right--
		}
		// 最右值小于枢轴，调换位置
		a[left], a[right] = a[right], a[left]
		// 最左值小于枢轴，左指针右移
		for left < right && a[left] < p {
			left++
		}
		// 最左值大于枢轴，调换位置
		a[left], a[right] = a[right], a[left]
	}
	// 返回最后指针（枢轴）位置
	return left
}
