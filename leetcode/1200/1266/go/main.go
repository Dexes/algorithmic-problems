package main

func minTimeToVisitAllPoints(points [][]int) int {
	result := 0
	for i := 1; i < len(points); i++ {
		x, y := abs(points[i-1][0]-points[i][0]), abs(points[i-1][1]-points[i][1])
		result += max(x, y)
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
