package main

func furthestDistanceFromOrigin(moves string) int {
	l, r := 0, 0
	for _, b := range moves {
		switch b {
		case 'L':
			l++
		case 'R':
			r++
		}
	}

	return len(moves) - l - r + abs(l-r)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
