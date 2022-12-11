package main

import "sort"

func deleteGreatestValue(grid [][]int) (result int) {
	for row := 0; row < len(grid); row++ {
		sort.Ints(grid[row])
	}

	for col := 0; col < len(grid[0]); col++ {
		result += max(grid, col)
	}

	return result
}

func max(grid [][]int, col int) int {
	result := grid[0][col]
	for row := 1; row < len(grid); row++ {
		if grid[row][col] > result {
			result = grid[row][col]
		}
	}

	return result
}
