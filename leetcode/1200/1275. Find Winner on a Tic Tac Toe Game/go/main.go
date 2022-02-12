package main

func tictactoe(moves [][]int) string {
	switch getWinner(makeMatrix(moves)) {
	case 'A':
		return "A"
	case 'B':
		return "B"
	case 'P':
		return "Pending"
	default:
		return "Draw"
	}
}

func makeMatrix(moves [][]int) [][]byte {
	matrix := [][]byte{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
	current, next := byte('A'), byte('B')
	for i := 0; i < len(moves); i++ {
		matrix[moves[i][0]][moves[i][1]] = current
		current, next = next, current
	}

	return matrix
}

func getWinner(matrix [][]byte) byte {
	if checkRow(matrix, 0) || checkCol(matrix, 0) {
		return matrix[0][0]
	}

	if checkRow(matrix, 2) || checkCol(matrix, 2) {
		return matrix[2][2]
	}

	if checkRow(matrix, 1) || checkCol(matrix, 1) || checkDiagonals(matrix) {
		return matrix[1][1]
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[i][j] < 'A' {
				return 'P'
			}
		}
	}

	return 'D'
}

func checkRow(matrix [][]byte, row int) bool {
	return matrix[row][0] == matrix[row][1] && matrix[row][0] == matrix[row][2]
}

func checkCol(matrix [][]byte, col int) bool {
	return matrix[0][col] == matrix[1][col] && matrix[0][col] == matrix[2][col]
}

func checkDiagonals(matrix [][]byte) bool {
	return matrix[1][1] == matrix[0][0] && matrix[1][1] == matrix[2][2] || matrix[1][1] == matrix[0][2] && matrix[1][1] == matrix[2][0]
}
