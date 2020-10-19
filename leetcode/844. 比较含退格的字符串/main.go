package main

import "fmt"

func main() {
	S := "ab#c"
	T := "acb#"
	res := backspaceCompare(S, T)
	fmt.Println(res)
}

// 简单思路：直接遍历重建字符串
func backspaceCompare(S string, T string) bool {
	return build(S) == build(T)
}

func build(str string) string {
	res := []byte{}
	for i := range str {
		if str[i] != '#' {
			res = append(res, str[i])
		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}
	fmt.Println(string(res))
	return string(res)
}
