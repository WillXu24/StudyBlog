package main

import (
	"fmt"
)

func main() {
	a := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	//bubbleSort(a)
	//selectSort(a)
	//insertSort(a)
	quickSort(a)
	fmt.Println(a)
	//sort.Ints(a)
	//fmt.Println(a)
}

/***** bubbleSort *****************************************************************************************************/
func bubbleSort(a []int) {
	hasReverse := true
	for i := 0; i < len(a)-1 && hasReverse; i++ {
		hasReverse = false
		for j := len(a) - 1; j > i; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				hasReverse = true
			}
		}
	}
}

/***** selectSort *****************************************************************************************************/
func selectSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}
}

/***** insertSort *****************************************************************************************************/
func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		index := i
		current := a[i]
		for j := i; j > 0 && a[j-1] > current; j-- {
			a[j] = a[j-1]
			index--
		}
		a[index] = current
	}
}

/***** quickSort *****************************************************************************************************/
func quickSort(a []int) {
	quickSortHandler(a, 0, len(a)-1)
}

func quickSortHandler(a []int, left, right int) {
	if left >= right {
		return
	}

	p, l, r := a[left], left, right
	for l < r {
		for l < r && a[r] >= p {
			r--
		}
		a[l], a[r] = a[r], a[l]
		for l < r && a[l] <= p {
			l++
		}
		a[l], a[r] = a[r], a[l]
	}

	quickSortHandler(a, left, l-1)
	quickSortHandler(a, l+1, right)
}

func partition(left, right int, a []int) int {
	pivot := a[left]
	for left < right {
		for left < right && a[right] >= pivot {
			right--
		}
		a[left], a[right] = a[right], a[left]
		for left < right && a[left] <= pivot {
			left++
		}
		a[left], a[right] = a[right], a[left]
	}
	return left
}
