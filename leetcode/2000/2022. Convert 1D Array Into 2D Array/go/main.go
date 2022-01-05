package main

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}

	result := makeMatrix(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = original[i*n+j]
		}
	}

	return result
}

func makeMatrix(m, n int) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	return result
}
