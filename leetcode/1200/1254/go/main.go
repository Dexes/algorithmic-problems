package main

func closedIsland(grid [][]int) (result int) {
	sinkEdges(grid)

	lastRow, lastCol := len(grid)-1, len(grid[0])-1
	for row := 1; row < lastRow; row++ {
		for col := 1; col < lastCol; col++ {
			if grid[row][col] == 0 {
				sink(grid, row, col)
				result++
			}
		}
	}

	return result
}

func sinkEdges(grid [][]int) {
	lastRow, lastCol := len(grid)-1, len(grid[0])-1

	for row := 0; row <= lastRow; row++ {
		if grid[row][0] == 0 {
			sink(grid, row, 0)
		}

		if grid[row][lastCol] == 0 {
			sink(grid, row, lastCol)
		}
	}

	for col := 0; col <= lastCol; col++ {
		if grid[0][col] == 0 {
			sink(grid, 0, col)
		}

		if grid[lastRow][col] == 0 {
			sink(grid, lastRow, col)
		}
	}
}

func sink(grid [][]int, row, col int) {
	if row < 0 || row == len(grid) || col < 0 || col == len(grid[0]) || grid[row][col] == 1 {
		return
	}

	grid[row][col] = 1
	sink(grid, row-1, col)
	sink(grid, row+1, col)
	sink(grid, row, col-1)
	sink(grid, row, col+1)
}
