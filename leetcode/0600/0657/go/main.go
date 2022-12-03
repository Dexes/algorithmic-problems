package main

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			x--
		case 'R':
			x++
		case 'U':
			y--
		case 'D':
			y++
		}
	}

	return x == 0 && y == 0
}
