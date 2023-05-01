package main

const land = 0

func findMaxFish(grid [][]int) (result int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == land {
				continue
			}

			x := sink(grid, i, j)
			if x > result {
				result = x
			}
		}
	}

	return result
}

func sink(grid [][]int, i, j int) int {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == land {
		return 0
	}

	current := grid[i][j]
	grid[i][j] = land

	return current +
		sink(grid, i-1, j) +
		sink(grid, i+1, j) +
		sink(grid, i, j-1) +
		sink(grid, i, j+1)
}
