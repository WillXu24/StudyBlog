package main

import "fmt"

func main() {
	a := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	res := ShellSort(a)
	fmt.Println(res)
}

// ShellSort 插入排序
func ShellSort(a []int) []int {
	// 定义间隔，采用2的次方缩小
	gap := 0
	for ; gap < len(a); gap = gap*2 + 1 {
	}
	// 间隔从大到小，直至为1
	for gap = (gap - 1) / 2; gap > 0; gap = (gap - 1) / 2 {
		// 从间隔值开始遍历，方式同插入排序
		for i := gap; i < len(a); i++ {
			temp := a[i]
			position := i
			// 注意此处循环命令不同
			for j := i; j-gap >= 0 && a[j-gap] > temp; j = j - gap {
				a[j] = a[j-gap]
				position = j - gap
			}
			a[position] = temp
		}
	}
	return a
}
