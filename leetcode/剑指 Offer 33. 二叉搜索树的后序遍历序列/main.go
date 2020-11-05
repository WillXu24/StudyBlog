package main

import "fmt"

func main() {
	//a := []int{1,6,3,2,5}
	a := []int{1, 3, 2, 6, 5}
	a = []int{4, 6, 7, 5}
	res := verifyPostorder(a)
	fmt.Println(res)
}

func verifyPostorder(postorder []int) bool {
	return help(postorder, 0, len(postorder)-1)
}

func help(a []int, left, root int) bool {
	// return
	if left >= root {
		return true
	}

	// find right
	right := left
	for ; right <= root && a[right] < a[root]; right++ {
	}
	//fmt.Println(right)

	// check right
	p := right
	for ; p <= root && a[p] > a[root]; p++ {
	}

	return (p == root) && help(a, left, right-1) && help(a, right, root-1)
}
