package main

import "fmt"

func main() {
	// name := "alex"
	// typed := "aaleex"
	// name := "saeed"
	// typed := "ssaaedd"
	// name := "leelee"
	// typed := "lleeelee"
	// name := "laiden"
	// typed := "laiden"
	name := "alex"
	typed := "alexxr"
	res := isLongPressedName(name, typed)
	fmt.Println(res)
}

func isLongPressedName(name string, typed string) bool {
	p1, p2 := 0, 0
	for p1 < len(name) && p2 < len(typed) {
		if name[p1] != typed[p2] {
			if p2 > 0 && typed[p2] == typed[p2-1] {
				p2++
				continue
			} else {
				return false
			}
		}
		p1++
		p2++
	}
	for p2 < len(typed) {
		if typed[p2] != typed[p2-1] {
			return false
		}
		p2++
	}
	return p1 == len(name)
}
