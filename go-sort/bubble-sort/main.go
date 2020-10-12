package main

import (
	"fmt"
)

func main() {
	a := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	res := BubbleSort(a)
	fmt.Println(res)
}

// BubbleSort 冒泡排序
func BubbleSort(a []int) []int {
	// 关键点2：设置标记
	flag := true
	for i := 0; i < len(a) && flag; i++ {
		flag = false
		// 关键点2：末尾开始
		for j := len(a) - 1; j > i; j-- {
			if a[j-1] > a[j] {
				a[j], a[j-1] = a[j-1], a[j]
				flag = true
			}
		}
	}
	return a
}
