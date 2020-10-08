package main

import "fmt"

func main() {
	fmt.Println(fib(5))
}

func fib(N int) int {
	if N < 2 {
		return N
	}
	v0, v1 := 0, 1
	for i := 0; i < N-1; i++ {
		v0, v1 = v1, v0+v1
	}
	return v1
}
