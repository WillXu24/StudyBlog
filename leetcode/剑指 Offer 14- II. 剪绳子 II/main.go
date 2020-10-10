package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(570))
}

// 不能用dp了!
// 不能用dp了!
// 不能用dp了!
// 因为取余数之后比较大小得到的结果是错误的，而先得到结果再取余数会溢出。。结果也是错误的
// # 贪心算法 # 数学解法
// 试试贪心
func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	if n == 4 {
		return 4
	}
	res := int64(1)
	for n > 4 {
		res *= 3
		res %= 1000000007
		n -= 3
	}
	//fmt.Println(res)
	return (int(res) * n) % 1000000007
}
