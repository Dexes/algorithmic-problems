package main

func countGoodRectangles(rectangles [][]int) int {
	squaresCount := make(map[int]int)
	maxSide := 0

	for i := 0; i < len(rectangles); i++ {
		side := min(rectangles[i][0], rectangles[i][1])
		maxSide = max(maxSide, side)
		squaresCount[side]++
	}

	return squaresCount[maxSide]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
