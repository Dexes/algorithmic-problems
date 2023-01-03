package main

type Vector2D struct {
	row, col int
	data     [][]int
}

func Constructor(vec [][]int) Vector2D {
	return Vector2D{row: 0, col: 0, data: vec}
}

func (x *Vector2D) Next() int {
	x.row, x.col = fixIndices(x.data, x.row, x.col)
	result := x.data[x.row][x.col]
	x.col++

	return result
}

func (x *Vector2D) HasNext() bool {
	if x.row < len(x.data) {
		x.row, x.col = fixIndices(x.data, x.row, x.col)
	}

	return x.row < len(x.data)
}

func fixIndices(data [][]int, row, col int) (int, int) {
	if col == len(data[row]) {
		for row++; row < len(data) && len(data[row]) == 0; row++ {
		}

		col = 0
	}

	return row, col
}
