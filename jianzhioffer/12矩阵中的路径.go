package main

var rows, cols int

func exist(board [][]byte, word string) bool {
	rows = len(board)
	cols = len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j< cols; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k int) bool {

	if i >= rows || i < 0 || j >= cols || j < 0 || board[i][j] != word[k] {
		return false
	}

	if k ==len(word) - 1 {
		return true
	}

	board[i][j] = 0
	res := dfs(board, word, i + 1, j, k + 1) || dfs(board, word, i - 1, j, k + 1) ||
		dfs(board, word, i, j + 1, k + 1) || dfs(board, word, i , j - 1, k + 1)
	board[i][j] = word[k]
	return res
}