package main

func generate(numRows int) [][]int {
	result := make([][]int, 0, numRows)
	result = append(result, []int{1}, []int{1, 1})

	for i := 2; i < numRows; i++ {
		current, prev := make([]int, i+1), result[i-1]
		current[0], current[i] = 1, 1
		for j := 1; j < i; j++ {
			current[j] = prev[j-1] + prev[j]
		}

		result = append(result, current)
	}

	return result[:numRows]
}
