package main

func findRotation(mat [][]int, target [][]int) bool {
	first, second, third, fourth := true, true, true, true
	l := len(mat) - 1

	for i := 0; i <= l; i++ {
		for j := 0; j <= l; j++ {
			first = first && mat[i][j] == target[i][j]
			second = second && mat[i][j] == target[j][l-i]
			third = third && mat[i][j] == target[l-i][l-j]
			fourth = fourth && mat[i][j] == target[l-j][i]

			if !(first || second || third || fourth) {
				return false
			}
		}
	}

	return true
}
