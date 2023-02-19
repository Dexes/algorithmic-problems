package main

const (
	water = 0
	land  = 1
)

var bytes = make([]byte, 0, 5000)

func numDistinctIslands(grid [][]int) int {
	islands := make(map[string]struct{})
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == land {
				hash := sink(0, grid, i, j, bytes)
				islands[string(hash[1:])] = struct{}{}
			}
		}
	}

	return len(islands)
}

func sink(direction byte, grid [][]int, row, col int, hash []byte) []byte {
	if row < 0 || row == len(grid) || col < 0 || col == len(grid[0]) || grid[row][col] == water {
		return hash
	}

	hash = append(hash, direction)
	grid[row][col] = water

	hash = sink('U', grid, row-1, col, hash)
	hash = sink('D', grid, row+1, col, hash)
	hash = sink('L', grid, row, col-1, hash)
	hash = sink('R', grid, row, col+1, hash)

	return append(hash, '|')
}
