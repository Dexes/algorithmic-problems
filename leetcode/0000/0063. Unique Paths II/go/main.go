package main

func uniquePathsWithObstacles(grid [][]int) int {
	m, n := len(grid)-1, len(grid[0])-1
	if grid[m][n] == 1 {
		return 0
	}

	prepareGrid(grid, m, n)
	fillGrid(grid, m, n)

	return -grid[m][n]
}

func prepareGrid(grid [][]int, m, n int) {
	for i := 0; i <= m && grid[i][0] != 1; i++ {
		grid[i][0] = -1
	}

	for i := 0; i <= n && grid[0][i] != 1; i++ {
		grid[0][i] = -1
	}
}

func fillGrid(grid [][]int, m, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i][j] == 1 {
				continue
			}

			if grid[i-1][j] != 1 {
				grid[i][j] += grid[i-1][j]
			}

			if grid[i][j-1] != 1 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
}
