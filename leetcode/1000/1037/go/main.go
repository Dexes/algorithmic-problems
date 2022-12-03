package main

func isBoomerang(points [][]int) bool {
	firstArea := (points[0][0] - points[1][0]) * (points[0][1] - points[2][1])
	secondArea := (points[0][0] - points[2][0]) * (points[0][1] - points[1][1])

	return firstArea != secondArea
}
