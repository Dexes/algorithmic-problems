package main

func findDiagonalOrder(mat [][]int) []int {
	m, n, i, j, direction := len(mat), len(mat[0]), 0, 0, 1
	result := make([]int, 0, m*n)

	for len(result) < cap(result) {
		for 0 <= i && i < m && 0 <= j && j < n {
			result = append(result, mat[i][j])
			i, j = i-direction, j+direction
		}

		i, j, direction = toNextDiagonal(m, n, i, j, direction)
	}

	return result
}

func toNextDiagonal(m, n, i, j, direction int) (int, int, int) {
	if direction == 1 {
		if j < n {
			return 0, j, -1
		}

		return i + 2, j - 1, -1
	}

	if i < m {
		return i, 0, 1
	}

	return i - 1, j + 2, 1
}
