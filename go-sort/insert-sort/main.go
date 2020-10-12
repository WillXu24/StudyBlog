package main

import "fmt"

func main() {
	a := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	res := InsertSort(a)
	fmt.Println(res)
}

// InsertSort 插入排序
func InsertSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		position := i
		for j := i; j > 0 && a[j-1] > temp; j-- {
			a[j] = a[j-1]
			position = j - 1
		}
		a[position] = temp
	}
	return a
}
