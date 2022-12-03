package main

func luckyNumbers(matrix [][]int) []int {
	result := make([]int, 0, len(matrix))
	for i := 0; i < len(matrix); i++ {
		x, col := min(matrix[i])
		if isMax(matrix, x, col) {
			result = append(result, x)
		}
	}

	return result
}

func min(a []int) (int, int) {
	x, col := a[0], 0
	for i := 1; i < len(a); i++ {
		if a[i] < x {
			x, col = a[i], i
		}
	}

	return x, col
}

func isMax(matrix [][]int, x, col int) bool {
	for i := 0; i < len(matrix); i++ {
		if x < matrix[i][col] {
			return false
		}
	}

	return true
}
