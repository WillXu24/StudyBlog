package main

import "fmt"

func main() {
	fmt.Println(fib(100))
}

// 递归超时
// func fib(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	return fib(n-1) + fib(n-2)
// }

// 动态规划
func fib(n int) int {
	if n < 2 {
		return n
	}
	v1, v2 := 0, 1
	for i := 1; i < n; i++ {
		v1, v2 = v2, (v1+v2)%1000000007
	}
	return v2 //% 1000000007
}
