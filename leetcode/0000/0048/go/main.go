package main

func rotate(matrix [][]int) {
	x, right := 0, len(matrix)-1
	for i := len(matrix)/2 - 1; i >= 0; i-- {
		for j := i; j < right-i; j++ {
			x, matrix[j][right-i] = matrix[j][right-i], matrix[i][j]
			x, matrix[right-i][right-j] = matrix[right-i][right-j], x
			x, matrix[right-j][i] = matrix[right-j][i], x
			matrix[i][j] = x
		}
	}
}
