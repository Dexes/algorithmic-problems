package main

func checkStraightLine(coordinates [][]int) bool {
	x, y := coordinates[0][0]-coordinates[1][0], coordinates[0][1]-coordinates[1][1]

	for i := 2; i < len(coordinates); i++ {
		firstArea := x * (coordinates[0][1] - coordinates[i][1])
		secondArea := y * (coordinates[0][0] - coordinates[i][0])

		if firstArea != secondArea {
			return false
		}
	}

	return true
}
