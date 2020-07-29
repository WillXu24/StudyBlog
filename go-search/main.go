package main

import "fmt"

func main() {
	// Linear Search
	fmt.Println(SequenceSearch([]int{9, 1, 5, 8, 3, 7, 4, 6, 2}, 2))

	fmt.Println(BinarySearch0([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4))
	fmt.Println(BinarySearch1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6))

	fmt.Println(InsertSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))

	fmt.Println(FibonacciSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9))
}
