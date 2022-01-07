package main

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)*len(matrix[0])
	for left < right {
		middle := (left + right) / 2
		i, j := middle/len(matrix[0]), middle%len(matrix[0])

		if target < matrix[i][j] {
			right = middle
			continue
		}

		if target > matrix[i][j] {
			left = middle + 1
			continue
		}

		return true
	}

	return false
}
