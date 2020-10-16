package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 5},
		{9, 8, 7, 6},
	}
	res := spiralOrder(matrix)
	fmt.Println(res)
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var res []int
	up, down, left, right, size := 0, len(matrix)-1, 0, len(matrix[0])-1, len(matrix)*len(matrix[0])
	for right >= left && len(res) < size {
		for i := left; i <= right && len(res) < size; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		for i := up; i <= down && len(res) < size; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		for i := right; i >= left && len(res) < size; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		for i := down; i >= up && len(res) < size; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
