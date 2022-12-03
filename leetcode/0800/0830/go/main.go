package main

func largeGroupPositions(s string) [][]int {
	result := make([][]int, 0, 100)
	for i, j := 0, 0; i < len(s); i = j {
		for j = i + 1; j < len(s) && s[i] == s[j]; j++ {
		}

		if j-i >= 3 {
			result = append(result, []int{i, j - 1})
		}
	}

	return result
}
