package main

func mostPoints(questions [][]int) int64 {
	for i := len(questions) - 2; i >= 0; i-- {
		nextIndex, next := i+questions[i][1]+1, 0
		if nextIndex < len(questions) {
			next = questions[nextIndex][0]
		}

		questions[i][0] = max(questions[i][0]+next, questions[i+1][0])
	}

	return int64(questions[0][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
