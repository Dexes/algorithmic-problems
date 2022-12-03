package main

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	return rectangleArea(ax1, ay1, ax2, ay2) +
		rectangleArea(bx1, by1, bx2, by2) -
		overlapArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
}

func overlapArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	x1, x2 := max(ax1, bx1), min(ax2, bx2)
	y1, y2 := max(ay1, by1), min(ay2, by2)

	if x1 >= x2 || y1 >= y2 {
		return 0
	}

	return rectangleArea(x1, y1, x2, y2)
}

func rectangleArea(x1, y1, x2, y2 int) int {
	return (x2 - x1) * (y2 - y1)
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
