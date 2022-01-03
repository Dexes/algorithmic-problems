package main

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	cols, rows := make([]int, 0, n), make([]int, 0, m)
	for col := 0; col < n; col++ {
		for row := 0; row < m; row++ {
			if matrix[col][row] == 0 {
				cols, rows = append(cols, col), append(rows, row)
			}
		}
	}

	for _, col := range cols {
		for row := 0; row < m; row++ {
			matrix[col][row] = 0
		}
	}

	for _, row := range rows {
		for col := 0; col < n; col++ {
			matrix[col][row] = 0
		}
	}
}
