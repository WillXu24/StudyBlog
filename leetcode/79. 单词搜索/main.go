package main

func main() {

}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k int) bool {
	if k == len(word) {
		return true
	}
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || word[k] != board[i][j] {
		return false
	}
	temp := board[i][j]
	board[i][j] = ' '
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i, j+1, k+1) ||
		dfs(board, word, i-1, j, k+1) || dfs(board, word, i, j-1, k+1)
	board[i][j] = temp
	return res
}
