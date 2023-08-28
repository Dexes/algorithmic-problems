package main

func candyCrush(board [][]int) [][]int {
	for repeat := true; repeat; repeat = crush(board) {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == 0 {
					continue
				}

				markLeft(board, i, j)
				markDown(board, i, j)
			}
		}
	}

	return board
}

func crush(board [][]int) (result bool) {
	for j := 0; j < len(board[0]); j++ {
		wIndex := len(board) - 1
		for i := wIndex; i >= 0; i-- {
			if board[i][j] < 0 {
				result = true
				continue
			}

			board[wIndex][j] = board[i][j]
			wIndex--
		}

		for ; wIndex >= 0; wIndex-- {
			board[wIndex][j] = 0
		}
	}

	return result
}

func markLeft(board [][]int, i, j int) {
	if j+2 >= len(board[i]) {
		return
	}

	x, y, z := abs(board[i][j]), abs(board[i][j+1]), abs(board[i][j+2])
	if x != y || x != z {
		return
	}

	negX := -x
	for ; j < len(board[i]) && abs(board[i][j]) == x; j++ {
		board[i][j] = negX
	}
}

func markDown(board [][]int, i, j int) {
	if i+2 >= len(board) {
		return
	}

	x, y, z := abs(board[i][j]), abs(board[i+1][j]), abs(board[i+2][j])
	if x != y || x != z {
		return
	}

	negX := -x
	for ; i < len(board) && abs(board[i][j]) == x; i++ {
		board[i][j] = negX
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
