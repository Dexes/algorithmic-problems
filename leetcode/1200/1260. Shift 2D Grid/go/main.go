package main

func shiftGrid(grid [][]int, k int) [][]int {
	size := len(grid) * len(grid[0])
	if k %= size; k == 0 {
		return grid
	}

	for counter, startIndex := 0, 0; counter < size; startIndex++ {
		value := get(grid, startIndex)
		for i := (startIndex + k) % size; i != startIndex; i = (i + k) % size {
			value = replace(grid, i, value)
			counter++
		}

		replace(grid, startIndex, value)
		counter++
	}

	return grid
}

func get(grid [][]int, index int) int {
	i, j := index/len(grid[0]), index%len(grid[0])
	return grid[i][j]
}

func replace(grid [][]int, index, value int) int {
	i, j := index/len(grid[0]), index%len(grid[0])
	value, grid[i][j] = grid[i][j], value

	return value
}
