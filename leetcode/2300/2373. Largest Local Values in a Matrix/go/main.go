package main

func largestLocal(grid [][]int) [][]int {
	rightIndex := len(grid) - 2
	for i := 0; i < rightIndex; i++ {
		for j := 0; j < rightIndex; j++ {
			setMax(grid, i, j)
		}

		grid[i] = grid[i][:rightIndex]
	}

	return grid[:rightIndex]
}

func setMax(grid [][]int, i, j int) {
	ks, ls := i+3, j+3
	if ks > len(grid) {
		ks = len(grid)
	}
	if ls > len(grid) {
		ls = len(grid)
	}

	for k := i; k < ks; k++ {
		for l := j; l < ls; j++ {
			if grid[k][l] > grid[i][j] {
				grid[i][j] = grid[k][l]
			}
		}
	}
}
