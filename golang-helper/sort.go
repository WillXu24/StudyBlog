package main

import "fmt"

func main() {
	a := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	//bubbleSort(a)
	//selectSort(a)
	insertSort(a)
	//quickSort(a)
	fmt.Println(a)
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
	quickSortHandler(0, len(a)-1, a)
}

func quickSortHandler(left, right int, a []int) {
	if left >= right {
		return
	}

	index := partition(left, right, a)
	quickSortHandler(left, index-1, a)
	quickSortHandler(index+1, right, a)
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
