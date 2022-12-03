package main

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		current, prev := triangle[i], triangle[i-1]
		current[0], current[i] = current[0]+prev[0], current[i]+prev[i-1]

		for j := 1; j < i; j++ {
			if prev[j-1] < prev[j] {
				current[j] += prev[j-1]
			} else {
				current[j] += prev[j]
			}
		}
	}

	last := triangle[len(triangle)-1]
	result := last[0]
	for i := 1; i < len(last); i++ {
		if last[i] < result {
			result = last[i]
		}
	}

	return result
}
