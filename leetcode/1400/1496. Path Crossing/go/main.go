package main

type point struct {
	x, y int
}

func isPathCrossing(path string) bool {
	current := point{0, 0}
	points := map[point]bool{current: true}
	for i := 0; i < len(path); i++ {
		current = nextPoint(current, path[i])
		if points[current] {
			return true
		}

		points[current] = true
	}

	return false
}

func nextPoint(p point, b byte) point {
	if b == 'N' {
		p.y++
		return p
	}

	if b == 'S' {
		p.y--
		return p
	}

	if b == 'E' {
		p.x++
		return p
	}

	p.x--
	return p
}
