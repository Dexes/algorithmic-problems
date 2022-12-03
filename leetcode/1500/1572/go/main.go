package main

func diagonalSum(mat [][]int) int {
	result, i, j := 0, 0, len(mat)-1
	for i < j {
		result += mat[i][i] + mat[i][j] + mat[j][i] + mat[j][j]
		i++
		j--
	}

	if len(mat)&1 == 1 {
		result += mat[i][i]
	}

	return result
}
