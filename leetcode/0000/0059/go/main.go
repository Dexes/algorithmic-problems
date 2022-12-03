package main

func generateMatrix(n int) [][]int {
	result := makeEmptyMatrix(n)
	max, current := n/2, 1
	for i := 0; i <= max; i++ {
		for j := i; j < n-i; j++ {
			result[i][j] = current
			current++
		}

		for j := i + 1; j < n-i-1; j++ {
			result[j][n-i-1] = current
			current++
		}

		for j := n - i - 1; j > i; j-- {
			result[n-i-1][j] = current
			current++
		}

		for j := n - i - 1; j > i; j-- {
			result[j][i] = current
			current++
		}
	}

	return result
}

func makeEmptyMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	return result
}
