package main

const (
	water = 0
	land  = 1
)

func numEnclaves(grid [][]int) (result int) {
	sinkEdges(grid)

	lastRow, lastCol := len(grid)-1, len(grid[0])-1
	for row := 1; row < lastRow; row++ {
		for col := 1; col < lastCol; col++ {
			result += grid[row][col]
		}
	}

	return result
}

func sinkEdges(grid [][]int) {
	lastRow, lastCol := len(grid)-1, len(grid[0])-1

	for row := 0; row <= lastRow; row++ {
		if grid[row][0] == land {
			sink(grid, row, 0)
		}

		if grid[row][lastCol] == land {
			sink(grid, row, lastCol)
		}
	}

	for col := 0; col <= lastCol; col++ {
		if grid[0][col] == land {
			sink(grid, 0, col)
		}

		if grid[lastRow][col] == land {
			sink(grid, lastRow, col)
		}
	}
}

func sink(grid [][]int, row, col int) {
	if row < 0 || row == len(grid) || col < 0 || col == len(grid[0]) || grid[row][col] == water {
		return
	}

	grid[row][col] = water
	sink(grid, row-1, col)
	sink(grid, row+1, col)
	sink(grid, row, col-1)
	sink(grid, row, col+1)
}
