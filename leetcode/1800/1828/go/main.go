package main

func countPoints(points [][]int, queries [][]int) []int {
	result := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		x, y, r := queries[i][0], queries[i][1], queries[i][2]*queries[i][2]
		for j := 0; j < len(points); j++ {
			if dx, dy := x-points[j][0], y-points[j][1]; dx*dx+dy*dy <= r {
				result[i]++
			}
		}
	}

	return result
}
