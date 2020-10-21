package main

import "fmt"

func main() {
	pushd := []int{1, 2, 3, 4, 5}
	//popped := []int{4, 5, 3, 2, 1}
	popped := []int{4, 3, 5, 1, 2}

	// pushd := []int{2, 1, 0}
	// popped := []int{1, 2, 0}

	res := validateStackSequences(pushd, popped)
	fmt.Println(res)
}

// #栈 20201021
// 思路：模拟栈
func validateStackSequences(pushed []int, popped []int) bool {
	var stark []int
	for i := 0; i < len(pushed); i++ {
		stark = append(stark, pushed[i])
		for len(stark) > 0 && stark[len(stark)-1] == popped[0] {
			stark = stark[:len(stark)-1]
			popped = popped[1:]
		}
	}
	return len(popped) == 0
}
