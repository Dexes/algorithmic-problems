package main

func largestTriangleArea(points [][]int) float64 {
	result := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				current := doubleArea(points[i], points[j], points[k])
				if current > result {
					result = current
				}
			}
		}
	}

	return float64(result) / 2
}

func doubleArea(p1, p2, p3 []int) int {
	result := p1[0]*(p2[1]-p3[1]) + p2[0]*(p3[1]-p1[1]) + p3[0]*(p1[1]-p2[1])
	if result < 0 {
		return -result
	}

	return result
}
