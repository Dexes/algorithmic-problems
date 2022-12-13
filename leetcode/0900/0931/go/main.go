package main

func minFallingPathSum(matrix [][]int) int {
	rightIndex := len(matrix) - 1
	if rightIndex == 0 {
		return matrix[0][0]
	}

	for row := 1; row <= rightIndex; row++ {
		prevRow := row - 1

		matrix[row][0] += min(matrix[prevRow][0], matrix[prevRow][1])
		matrix[row][rightIndex] += min(matrix[prevRow][rightIndex-1], matrix[prevRow][rightIndex])

		for col := rightIndex - 1; col > 0; col-- {
			matrix[row][col] += min(min(matrix[prevRow][col-1], matrix[prevRow][col]), matrix[prevRow][col+1])
		}
	}

	result := matrix[rightIndex][0]
	for i := 1; i <= rightIndex; i++ {
		result = min(result, matrix[rightIndex][i])
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
