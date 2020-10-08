package main

import "fmt"

func main() {
	fmt.Println(numWays(7))
}

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	n0, n1 := 1, 1
	for i := 0; i < n-1; i++ {
		n0, n1 = n1, (n0+n1)%1000000007
	}
	return n1 % 1000000007
}
