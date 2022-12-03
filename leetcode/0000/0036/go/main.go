package main

func isValidSudoku(board [][]byte) bool {
	rows, columns, squares := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board[row][column] == '.' {
				continue
			}

			num, square := board[row][column]-'1', row/3*3+column/3
			if rows[row][num] || columns[column][num] || squares[square][num] {
				return false
			}

			rows[row][num], columns[column][num], squares[square][num] = true, true, true
		}
	}

	return true
}
