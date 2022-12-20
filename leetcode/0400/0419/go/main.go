package main

func countBattleships(board [][]byte) (result int) {
	rows, cols := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'X' {
				remove(i, j, board)
				result++
			}
		}
	}

	return result
}

func remove(i, j int, board [][]byte) {
	for index := i + 1; index < len(board) && board[index][j] == 'X'; index++ {
		board[index][j] = '.'
	}

	for index := j + 1; index < len(board[i]) && board[i][index] == 'X'; index++ {
		board[i][index] = '.'
	}
}
