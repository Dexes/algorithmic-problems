package main

func minFlips(a int, b int, c int) (result int) {
	for ; a > 0 || b > 0 || c > 0; a, b, c = a>>1, b>>1, c>>1 {
		switch x, y, z := a&1, b&1, c&1; {
		case x == 1 && y == 1 && z == 0:
			result += 2
		case (x == 1 || y == 1) && z == 0:
			result += 1
		case x == 0 && y == 0 && z == 1:
			result += 1
		}
	}

	return result
}
