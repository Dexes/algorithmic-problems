package main

func mergeArrays(x, y [][]int) [][]int {
	result := make([][]int, 0, len(x)+len(y))

	i, j := 0, 0
	for i < len(x) && j < len(y) {
		switch {
		case x[i][0] < y[j][0]:
			result = append(result, x[i])
			i++
		case x[i][0] > y[j][0]:
			result = append(result, y[j])
			j++
		default:
			x[i][1] += y[j][1]
			result = append(result, x[i])
			i, j = i+1, j+1
		}
	}

	return append(append(result, x[i:]...), y[j:]...)
}
