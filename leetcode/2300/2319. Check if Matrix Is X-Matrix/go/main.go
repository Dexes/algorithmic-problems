package main

func checkXMatrix(grid [][]int) bool {
	var (
		rightIndex         = len(grid) - 1
		isDiagonal, isZero bool
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if isDiagonal, isZero = i == j || i == rightIndex-j, grid[i][j] == 0; isDiagonal == isZero {
				return false
			}
		}
	}

	return true
}
