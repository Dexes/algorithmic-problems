package main

func searchMatrix(matrix [][]int, target int) bool {
	for row, column := 0, len(matrix[0])-1; row < len(matrix) && column >= 0; {
		switch {
		case matrix[row][column] < target:
			row++
		case matrix[row][column] > target:
			column--
		default:
			return true
		}
	}

	return false
}
