package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if check(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func check(board [][]byte, word string, i, j, x int) bool {
	if x == len(word) {
		return true
	}

	if i < 0 || j < 0 || i == len(board) || j == len(board[i]) || board[i][j] != word[x] {
		return false
	}

	b := board[i][j]
	board[i][j] = 0

	result := check(board, word, i-1, j, x+1) ||
		check(board, word, i+1, j, x+1) ||
		check(board, word, i, j-1, x+1) ||
		check(board, word, i, j+1, x+1)

	board[i][j] = b
	return result
}
