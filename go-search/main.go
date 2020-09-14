package main

import "fmt"

func main() {
	//fmt.Println(SequenceSearch([]int{9, 1, 5, 8, 3, 7, 4, 6, 2}, 2))
	//
	//fmt.Println(BinarySearch0([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4))
	fmt.Println(BinarySearch1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0))
	//
	fmt.Println(InsertSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8))

	//fmt.Println(FibonacciSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9))
}

/**********************************************************************************************************************/
func SequenceSearch(a []int, b int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return i
		}
	}
	return -1
}

	fmt.Println(InsertSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))

func bSearch(a []int, l, h, b int) int {
	if l>h{
		return -1
	}
	m := (l + h) / 2
	if a[m] == b {
		return m
	}
	if a[m] > b {
		return bSearch(a, l, m-1, b)
	}
	if a[m] < b {
		return bSearch(a, m+1, h, b)
	}
	return -1
}

/**********************************************************************************************************************/
func InsertSearch(a []int, b int) int {
	l, m, h := 0, 0, len(a)-1
	for l <= h {
		m = l + (b-a[l])/(a[h]-a[l])*(h-l)
		if a[m] == b {
			return m
		}
		if a[m] > b {
			h = m - 1
		}
		if a[m] < b {
			l = m + 1
		}
	}
	return -1
}
