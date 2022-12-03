package main

func projectionArea(grid [][]int) int {
	result, n := 0, len(grid)
	for i := 0; i < n; i++ {
		x, y := 0, 0

		for j := 0; j < n; j++ {
			x, y = max(x, grid[i][j]), max(y, grid[j][i])
			if grid[i][j] > 0 {
				result++
			}
		}

		result += x + y
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
