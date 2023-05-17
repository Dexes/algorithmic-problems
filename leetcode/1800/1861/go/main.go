package main

const (
	stone    = '#'
	obstacle = '*'
	empty    = '.'
)

func rotateTheBox(box [][]byte) [][]byte {
	result := makeBox(len(box[0]), len(box))
	for row := 0; row < len(box); row++ {
		processRow(row, box, result)
	}

	return result
}

func processRow(boxRow int, box, result [][]byte) {
	boxCol := len(box[boxRow]) - 1
	resultRow, resultCol := boxCol, len(box)-boxRow-1
	for ; boxCol >= 0; boxCol-- {
		if box[boxRow][boxCol] == obstacle {
			setValues(result, resultCol, boxCol+1, resultRow, empty)
			result[boxCol][resultCol] = obstacle
			resultRow = boxCol - 1
			continue
		}

		if box[boxRow][boxCol] == stone {
			result[resultRow][resultCol] = stone
			resultRow--
		}
	}

	setValues(result, resultCol, 0, resultRow, empty)
}

func setValues(data [][]byte, col, rowStart, rowEnd int, value byte) {
	for ; rowStart <= rowEnd; rowStart++ {
		data[rowStart][col] = value
	}
}

func makeBox(n, m int) [][]byte {
	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = make([]byte, m)
	}

	return result
}
