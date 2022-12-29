package main

import "sort"

type BinaryMatrix struct {
	Get        func(int, int) int
	Dimensions func() []int
}

func leftMostColumnWithOne(matrix BinaryMatrix) int {
	dimensions := matrix.Dimensions()
	rows, cols := dimensions[0], dimensions[1]
	col := cols

	for row := 0; row < rows; row++ {
		leftCol := sort.Search(col, func(i int) bool { return matrix.Get(row, i) == 1 })
		if leftCol < col {
			col = leftCol
		}
	}

	if col == cols {
		return -1
	}

	return col
}
