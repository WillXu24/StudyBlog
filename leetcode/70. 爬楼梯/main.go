package main

import "fmt"

func main() {
	res := climbStairs(5)
	fmt.Println(res)
}

func climbStairs(n int) int {
	a := []int{1, 1}
	for i := 2; i <= n; i++ {
		a = append(a, a[i-1]+a[i-2])
	}
	return a[n]
}
