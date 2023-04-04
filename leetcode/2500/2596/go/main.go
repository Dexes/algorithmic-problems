package main

var shifts = [][]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}

func checkValidGrid(grid [][]int) bool {
	if len(grid) < 5 {
		return false
	}

	row, col, success, max := 0, 0, true, len(grid)*len(grid)
	for step := 1; success && step < max; step++ {
		row, col, success = next(grid, row, col, step)
	}

	return success
}

func next(grid [][]int, row, col, step int) (int, int, bool) {
	for _, shift := range shifts {
		if r, c := row+shift[0], col+shift[1]; isValid(len(grid), r, c) && grid[r][c] == step {
			return r, c, true
		}
	}

	return 0, 0, false
}

func isValid(max, row, col int) bool {
	return 0 <= row && row < max && 0 <= col && col < max
}
