package main

type SubrectangleQueries struct {
	Data [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{Data: rectangle}
}

func (x *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			x.Data[i][j] = newValue
		}
	}
}

func (x *SubrectangleQueries) GetValue(row int, col int) int {
	return x.Data[row][col]
}
