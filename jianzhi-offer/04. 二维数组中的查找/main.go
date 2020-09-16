package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	matrix = [][]int{{5}}
	res := findNumberIn2DArray(matrix, 5)
	fmt.Println(res)
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 排除空矩阵
	row := len(matrix)
	if row == 0 {
		//fmt.Println("row",row)
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		//fmt.Println("col",col)
		return false
	}

	for i := col - 1; i >= 0; i-- {
		//fmt.Println(matrix[0][i])
		// 注意是小于等于
		if matrix[0][i] <= target {
			for j := 0; j < row; j++ {
				if matrix[j][i] == target {
					return true
				}
			}
		}
	}

	//fmt.Println(matrix[0][col-1])
	//fmt.Println(row,col)
	return false
}
