package main

import "fmt"

func main() {
	fmt.Println(myPow(2, -1))
}

func myPow(x float64, n int) float64 {
	var res float64
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}

	temp := myPow(x, n/2)
	res = temp * temp

	if n%2 == 1 {
		res *= x
	}
	if n%2 == -1 {
		res *= 1 / x
	}

	return res
}
