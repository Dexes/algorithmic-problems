package main

func surfaceArea(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			result += area(grid, i, j)
		}
	}

	return result
}

func area(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}

	result := grid[i][j]*4 + 2

	if i > 0 {
		result -= min(grid[i][j], grid[i-1][j]) * 2
	}

	if j > 0 {
		result -= min(grid[i][j], grid[i][j-1]) * 2
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
