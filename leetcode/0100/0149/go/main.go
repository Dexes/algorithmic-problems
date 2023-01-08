package main

func maxPoints(points [][]int) (result int) {
	if len(points) < 3 {
		return len(points)
	}

	for i := 0; i < len(points); i++ {
		result = max(result, lines(points, i))
	}

	return result + 1
}

func lines(points [][]int, current int) map[float64]int {
	result := make(map[float64]int)
	for i := 0; i < len(points); i++ {
		if current == i {
			continue
		}

		angle := float64(points[i][0]-points[current][0]) / float64(points[i][1]-points[current][1])
		result[angle]++
	}

	return result
}

func max(current int, lines map[float64]int) int {
	for _, x := range lines {
		if x > current {
			current = x
		}
	}

	return current
}
