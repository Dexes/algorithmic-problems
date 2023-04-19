package main

func findColumnWidth(grid [][]int) []int {
	for col := 0; col < len(grid[0]); col++ {
		grid[0][col] = width(longest(grid, col))
	}

	return grid[0]
}

func width(x int) (result int) {
	for ; x > 0; x /= 10 {
		result++
	}

	return result
}

func longest(grid [][]int, col int) int {
	max := 1
	for i := 0; i < len(grid); i++ {
		if grid[i][col] < 0 {
			grid[i][col] *= -10
		}

		if grid[i][col] > max {
			max = grid[i][col]
		}
	}

	return max
}
