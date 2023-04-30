package main

type Coordinates struct {
	Row int
	Col int
}

func firstCompleteIndex(steps []int, matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	if m == 1 || n == 1 {
		return 0
	}

	indices := make([]Coordinates, n*m+1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			indices[matrix[i][j]].Row = i
			indices[matrix[i][j]].Col = j
		}
	}

	rows := make([]int, n)
	cols := make([]int, m)

	for i := 0; ; i++ {
		coordinates := indices[steps[i]]

		if rows[coordinates.Row]++; rows[coordinates.Row] == m {
			return i
		}

		if cols[coordinates.Col]++; cols[coordinates.Col] == n {
			return i
		}
	}
}
