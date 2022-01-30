package main

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	if m == 1 && n == 1 {
		return img
	}

	if m == 1 {
		return smoothRow(img, n)
	}

	if n == 1 {
		return smoothColumn(img, m)
	}

	result := allocate(m, n)
	for i := m - 2; i > 0; i-- {
		for j := n - 2; j > 0; j-- {
			result[i][j] = img[i-1][j-1] + img[i][j-1] + img[i+1][j-1] +
				img[i-1][j] + img[i][j] + img[i+1][j] +
				img[i-1][j+1] + img[i][j+1] + img[i+1][j+1]

			result[i][j] /= 9
		}
	}

	for i := m - 2; i > 0; i-- {
		result[i][0] = (img[i-1][0] + img[i][0] + img[i+1][0] + img[i-1][1] + img[i][1] + img[i+1][1]) / 6
		result[i][n-1] = (img[i-1][n-1] + img[i][n-1] + img[i+1][n-1] + img[i-1][n-2] + img[i][n-2] + img[i+1][n-2]) / 6
	}

	for i := n - 2; i > 0; i-- {
		result[0][i] = (img[0][i-1] + img[0][i] + img[0][i+1] + img[1][i-1] + img[1][i] + img[1][i+1]) / 6
		result[m-1][i] = (img[m-1][i-1] + img[m-1][i] + img[m-1][i+1] + img[m-2][i-1] + img[m-2][i] + img[m-2][i+1]) / 6
	}

	result[0][0] = (img[0][0] + img[0][1] + img[1][0] + img[1][1]) / 4
	result[0][n-1] = (img[0][n-1] + img[0][n-2] + img[1][n-1] + img[1][n-2]) / 4
	result[m-1][0] = (img[m-1][0] + img[m-1][1] + img[m-2][0] + img[m-2][1]) / 4
	result[m-1][n-1] = (img[m-1][n-1] + img[m-1][n-2] + img[m-2][n-1] + img[m-2][n-2]) / 4

	return result
}

func smoothRow(img [][]int, n int) [][]int {
	result := allocate(1, n)
	result[0][0] = (img[0][0] + img[0][1]) / 2
	result[0][n-1] = (img[0][n-1] + img[0][n-2]) / 2
	for i := n - 2; i > 0; i-- {
		result[0][i] = (img[0][i-1] + img[0][i] + img[0][i+1]) / 3
	}

	return result
}

func smoothColumn(img [][]int, m int) [][]int {
	result := allocate(m, 1)
	result[0][0] = (img[0][0] + img[1][0]) / 2
	result[m-1][0] = (img[m-1][0] + img[m-2][0]) / 2
	for i := m - 2; i > 0; i-- {
		result[i][0] = (img[i-1][0] + img[i][0] + img[i+1][0]) / 3
	}

	return result
}

func allocate(m, n int) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	return result
}
