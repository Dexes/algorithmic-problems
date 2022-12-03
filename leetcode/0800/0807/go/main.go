package main

func maxIncreaseKeepingSkyline(grid [][]int) (result int) {
	rows, columns := maximums(grid)
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid); column++ {
			result += min(rows[row], columns[column]) - grid[row][column]
		}
	}

	return result
}

func maximums(grid [][]int) ([]int, []int) {
	rows, columns := make([]int, len(grid)), make([]int, len(grid))
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid); column++ {
			if grid[row][column] > rows[row] {
				rows[row] = grid[row][column]
			}

			if grid[row][column] > columns[column] {
				columns[column] = grid[row][column]
			}
		}
	}

	return rows, columns
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
