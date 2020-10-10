package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(10))
	// fmt.Println(cuttingRope(57))
	// fmt.Println(cuttingRope(999)%1000000007)
	// for i:=50;i<70;i++{
	// 	if cuttingRope(i)>1000000007{
	// 		fmt.Println(i)
	// 	}
	// }
}

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), dp[i-j]*j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
