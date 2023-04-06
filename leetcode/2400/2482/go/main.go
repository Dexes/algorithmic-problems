package main

func onesMinusZeros(grid [][]int) [][]int {
	rows, cols := make([]int, len(grid)), make([]int, len(grid[0]))
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 0 {
				rows[r], cols[c] = rows[r]-1, cols[c]-1
			} else {
				rows[r], cols[c] = rows[r]+1, cols[c]+1
			}
		}
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			grid[r][c] = rows[r] + cols[c]
		}
	}

	return grid
}
