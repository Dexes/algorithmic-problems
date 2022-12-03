package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	x1, x2 := max(rec1[0], rec2[0]), min(rec1[2], rec2[2])
	y1, y2 := max(rec1[1], rec2[1]), min(rec1[3], rec2[3])

	return x2 > x1 && y2 > y1
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
