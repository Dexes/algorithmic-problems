package main

func solve(board [][]byte) {
	lastRow, lastCol := len(board)-1, len(board[0])-1
	if lastRow < 2 || lastCol < 2 {
		return
	}

	blocked := blockEdges(board)
	for r := 1; r < lastRow; r++ {
		for c := 1; c < lastCol; c++ {
			if !blocked[r][c] {
				board[r][c] = 'X'
			}
		}
	}
}

func blockEdges(board [][]byte) [][]bool {
	lastRow, lastCol := len(board)-1, len(board[0])-1
	blocked := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		blocked[i] = make([]bool, len(board[i]))
	}

	for r := 0; r < len(board); r++ {
		if !blocked[r][0] && board[r][0] == 'O' {
			blockDFS(board, blocked, r, 0)
		}

		if !blocked[r][lastCol] && board[r][lastCol] == 'O' {
			blockDFS(board, blocked, r, lastCol)
		}
	}

	for c := 0; c < len(board[0]); c++ {
		if !blocked[0][c] && board[0][c] == 'O' {
			blockDFS(board, blocked, 0, c)
		}

		if !blocked[lastRow][c] && board[lastRow][c] == 'O' {
			blockDFS(board, blocked, lastRow, c)
		}
	}

	return blocked
}

func blockDFS(board [][]byte, blocked [][]bool, r, c int) {
	if r < 0 || r == len(board) || c < 0 || c == len(board[r]) || blocked[r][c] || board[r][c] == 'X' {
		return
	}

	blocked[r][c] = true
	blockDFS(board, blocked, r-1, c)
	blockDFS(board, blocked, r+1, c)
	blockDFS(board, blocked, r, c-1)
	blockDFS(board, blocked, r, c+1)
}
