package main

func uniquePaths(m, n int) int {
	grid := makeGrid(m, n)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]
}

func makeGrid(m, n int) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		result[i][0] = 1
	}

	for i := 0; i < n; i++ {
		result[0][i] = 1
	}

	return result
}
