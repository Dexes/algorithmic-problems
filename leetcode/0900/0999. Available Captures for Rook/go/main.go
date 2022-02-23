package main

func numRookCaptures(board [][]byte) int {
	i, j := findRook(board)

	return verticalCapture(board, i, j, 1) +
		verticalCapture(board, i, j, -1) +
		horizontalCapture(board, i, j, 1) +
		horizontalCapture(board, i, j, -1)
}

func verticalCapture(board [][]byte, i, j, step int) int {
	for i += step; 0 <= i && i < 8; i += step {
		if board[i][j] == 'p' {
			return 1
		}

		if board[i][j] == 'B' {
			return 0
		}
	}

	return 0
}

func horizontalCapture(board [][]byte, i, j, step int) int {
	for j += step; 0 <= j && j < 8; j += step {
		if board[i][j] == 'p' {
			return 1
		}

		if board[i][j] == 'B' {
			return 0
		}
	}

	return 0
}

func findRook(board [][]byte) (int, int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				return i, j
			}
		}
	}

	return -1, -1
}
