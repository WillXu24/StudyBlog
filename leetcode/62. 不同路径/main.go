package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

// #动态规划 #多维数组
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了30.35%的用户
func uniquePaths(m int, n int) int {
	// 初始化二维数组
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
		paths[i][0] = 1
		for j := 1; j < n; j++ {
			if i == 0 {
				paths[0][j] = 1
			} else {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}

		}
	}
	//// 多点重复计算
	//for i := 1; i < m; i++ {
	//	for j := 1; j < n; j++ {
	//
	//	}
	//}
	return paths[m-1][n-1]
}
