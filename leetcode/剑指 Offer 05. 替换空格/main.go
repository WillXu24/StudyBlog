package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "We are happy."
	res := replaceSpace(str)
	fmt.Println(res)
}

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
