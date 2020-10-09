package main

import "fmt"

func main() {
	fmt.Println(movingCount(2, 3, 1))
}

// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标
// [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格
// 外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器
// 人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因
// 为3+5+3+8=19。请问该机器人能够到达多少个格子？
func movingCount(m int, n int, k int) int {
	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, n)
	}
	dfs(b, m, n, k, 0, 0)
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += b[i][j]
		}
	}
	return res
}

func dfs(b [][]int, m, n, k, i, j int) {
	if i >= m || j >= n || j < 0 || i < 0 {
		return
	}
	if i/10+i%10+j/10+j%10 > k || b[i][j] != 0 {
		return
	}
	b[i][j] = 1
	dfs(b, m, n, k, i+1, j)
	dfs(b, m, n, k, i, j+1)
	dfs(b, m, n, k, i-1, j)
	dfs(b, m, n, k, i, j-1)
}
