package main

func solveSudoku(board [][]byte) {
	NewSolver(board).Solve()
}

type Solver struct {
	board                  [][]byte
	rows, columns, squares [9][9]bool
}

func NewSolver(board [][]byte) *Solver {
	result := &Solver{board: board}
	result.makeRCS(board)

	return result
}

func (s *Solver) Solve() {
	s.solve(0, 0)
}

func (s *Solver) solve(row, column int) bool {
	row, column = s.nextEmpty(row, column)
	if row == -1 || column == -1 {
		return true
	}

	square := row/3*3 + column/3

	for num := 0; num < 9; num++ {
		if s.rows[row][num] || s.columns[column][num] || s.squares[square][num] {
			continue
		}

		s.rows[row][num], s.columns[column][num], s.squares[square][num] = true, true, true
		s.board[row][column] = byte(num + '1')

		if s.solve(row, column+1) {
			return true
		}

		s.rows[row][num], s.columns[column][num], s.squares[square][num] = false, false, false
	}

	s.board[row][column] = '.'
	return false
}

func (s *Solver) nextEmpty(row, column int) (int, int) {
	board := s.board
	for ; column < 9; column++ {
		if board[row][column] == '.' {
			return row, column
		}
	}

	for row++; row < 9; row++ {
		for column = 0; column < 9; column++ {
			if board[row][column] == '.' {
				return row, column
			}
		}
	}

	return -1, -1
}

func (s *Solver) makeRCS(board [][]byte) {
	rows, columns, squares := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board[row][column] == '.' {
				continue
			}

			num, square := board[row][column]-'1', row/3*3+column/3
			rows[row][num], columns[column][num], squares[square][num] = true, true, true
		}
	}

	s.rows, s.columns, s.squares = rows, columns, squares
}
