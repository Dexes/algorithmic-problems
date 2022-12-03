package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if c*r != len(mat)*len(mat[0]) || r == len(mat) && c == len(mat[0]) {
		return mat
	}

	result := makeMatrix(r, c)
	r, c = 0, 0

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			result[r][c] = mat[i][j]
			c++

			if c == len(result[r]) {
				r, c = r+1, 0
			}
		}
	}

	return result
}

func makeMatrix(r, c int) [][]int {
	result := make([][]int, c)
	for i := 0; i < c; i++ {
		result[i] = make([]int, r)
	}

	return result
}
