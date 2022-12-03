package main

func findDiagonalOrder(nums [][]int) []int {
	diagonals, diagonal, capacity := make([][]int, getLastDiagonal(nums)), 0, 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if diagonal = i + j; diagonals[diagonal] == nil {
				diagonals[diagonal] = make([]int, 0, 8)
			}

			diagonals[diagonal] = append(diagonals[diagonal], nums[i][j])
		}

		capacity += len(nums[i])
	}

	result := make([]int, 0, capacity)
	for diagonal = 0; diagonal < len(diagonals); diagonal++ {
		for i := len(diagonals[diagonal]) - 1; i >= 0; i-- {
			result = append(result, diagonals[diagonal][i])
		}
	}

	return result
}

func getLastDiagonal(nums [][]int) (result int) {
	for i := 0; i < len(nums); i++ {
		if diagonal := i + len(nums[i]); diagonal > result {
			result = diagonal
		}
	}

	return result
}
