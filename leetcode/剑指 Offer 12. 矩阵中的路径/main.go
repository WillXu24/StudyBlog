package main

import "fmt"

func main() {
	// b := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// word := "ABCCED"
	b := [][]byte{
		{'a', 'b'},
	}
	word := "ba"
	fmt.Println(exist(b, word))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0, word, board) {
				return true
			}
		}
	}
	return false
}

func dfs(i, j, k int, word string, board [][]byte) bool {
	if k == len(word) {
		return true
	}
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || word[k] != board[i][j] {
		//fmt.Println("err")
		return false
	}
	// if k == len(word)-1 {
	// 	return true
	// }
	temp := board[i][j]
	board[i][j] = ' '
	res := dfs(i+1, j, k+1, word, board) || dfs(i, j+1, k+1, word, board) ||
		dfs(i-1, j, k+1, word, board) || dfs(i, j-1, k+1, word, board)
	//fmt.Println(res)
	board[i][j] = temp
	return res
}
