package main

func countSubIslands(first, second [][]int) (result int) {
	for row := 0; row < len(second); row++ {
		for col := 0; col < len(second[row]); col++ {
			if second[row][col] == 1 && isSub(first, second, row, col) {
				result++
			}
		}
	}

	return result
}

func isSub(first, second [][]int, row, col int) bool {
	if row < 0 || row == len(second) || col < 0 || col == len(second[0]) || second[row][col] == 0 {
		return true
	}

	result := first[row][col] == 1
	first[row][col], second[row][col] = 0, 0

	result = isSub(first, second, row-1, col) && result
	result = isSub(first, second, row+1, col) && result
	result = isSub(first, second, row, col-1) && result
	return isSub(first, second, row, col+1) && result
}
