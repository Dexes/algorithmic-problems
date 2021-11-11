package main

func spiralOrder(matrix [][]int) []int {
	rows, columns := len(matrix), len(matrix[0])
	result := make([]int, 0, rows*columns)

	for loop := 0; len(result) < cap(result); loop++ {
		for column := loop; column < columns-loop && len(result) < cap(result); column++ {
			result = append(result, matrix[loop][column])
		}

		for row := loop + 1; row < rows-loop-1 && len(result) < cap(result); row++ {
			result = append(result, matrix[row][columns-loop-1])
		}

		for column := columns - loop - 1; column > loop && len(result) < cap(result); column-- {
			result = append(result, matrix[rows-loop-1][column])
		}

		for row := rows - loop - 1; row > loop && len(result) < cap(result); row-- {
			result = append(result, matrix[row][loop])
		}
	}

	return result
}
