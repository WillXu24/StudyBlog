package main

import "fmt"

func main() {
	A := []int{-4, -1, 0, 3, 10}
	res := sortedSquares(A)
	fmt.Println(res)
}

func sortedSquares(A []int) []int {
	length := len(A)
	res := make([]int, length)
	start, end := 0, length-1
	for i := end; i >= 0; i-- {
		a, b := A[start]*A[start], A[end]*A[end]
		if a > b {
			res[i] = a
			start++
		} else {
			res[i] = b
			end--
		}
	}
	return res
}
