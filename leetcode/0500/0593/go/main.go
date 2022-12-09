package main

func validSquare(p1, p2, p3, p4 []int) bool {
	diagonal, a, b, c, d := distance(p1, p2), p1, p3, p2, p4
	if diagonal == 0 {
		return false
	}

	if length := distance(p1, p3); length > diagonal {
		diagonal, a, b, c, d = length, p1, p2, p3, p4
	}

	if length := distance(p1, p4); length > diagonal {
		diagonal, a, b, c, d = length, p1, p2, p4, p3
	}

	side := distance(a, b)
	return distance(b, c) == side && distance(c, d) == side && distance(b, d) == diagonal
}

func distance(a, b []int) int {
	x, y := a[0]-b[0], a[1]-b[1]
	return x*x + y*y
}
