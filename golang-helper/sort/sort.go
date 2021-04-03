package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	//bubbleSort(arr)
	//selectSort(arr)
	//insertSort(arr)
	//shellSort(arr)
	//arr = mergeSort(arr)
	//quickSort(arr)
	heapSort(arr)
	fmt.Println(arr)
}

/***** bubbleSort *****************************************************************************************************/
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

/***** selectSort *****************************************************************************************************/
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

/***** insertSort *****************************************************************************************************/
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > cur; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = cur
	}
}

/***** shellSort *****************************************************************************************************/
func shellSort(arr []int) {
	gap := 1
	for ; gap < len(arr)-1; gap = gap*2 + 1 {
	}

	for ; gap > 0; gap = gap / 2 {
		for i := 0; i < len(arr); i++ {
			cur := arr[i]
			j := i
			for ; j-gap >= 0 && arr[j-gap] > cur; j = j - gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = cur
		}
	}
}

/***** mergeSort *****************************************************************************************************/
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	middle := len(arr) / 2
	return mergeHelper(mergeSort(arr[0:middle]), mergeSort(arr[middle:]))
}

func mergeHelper(a, b []int) []int {
	var res []int

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
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

/***** quickSort *****************************************************************************************************/
func quickSort(arr []int) {
	quickSortHandler(arr, 0, len(arr)-1)
}

func quickSortHandler(arr []int, left, right int) {
	if left >= right {
		return
	}

	p, l, r := arr[left], left, right
	for l < r {
		for l < r && arr[r] >= p {
			r--
		}
		arr[l], arr[r] = arr[r], arr[l]
		for l < r && arr[l] <= p {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}

	quickSortHandler(arr, left, l-1)
	quickSortHandler(arr, l+1, right)
}

/***** quickSort *****************************************************************************************************/
func heapSort(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		heapSink(arr, i)
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapSink(arr[:i], 0)
	}
}

func heapSink(arr []int, root int) {
	for i := root*2 + 1; i < len(arr); i = i*2 + 1 {
		if i+1 < len(arr) && arr[i] < arr[i+1] {
			i++
		}
		if arr[root] > arr[i] {
			break
		}
		arr[root], arr[i], root = arr[i], arr[root], i
	}
}
