package main

func nearestValidPoint(x int, y int, points [][]int) int {
	result, minDistance := -1, 20000
	for i := 0; i < len(points); i++ {
		if points[i][0] != x && points[i][1] != y {
			continue
		}

		distance := sub(points[i][0], x) + sub(points[i][1], y)
		if distance < minDistance {
			result, minDistance = i, distance
		}
	}

	return result
}

func sub(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}
