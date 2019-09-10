package main

import (
	"fmt"
	"strings"
)

func main()  {
	str06 := strings.TrimSpace(" hello hao hao hao ")
	fmt.Println("删除开头末尾的空格 =" + str06) //'hello hao hao hao'
}
