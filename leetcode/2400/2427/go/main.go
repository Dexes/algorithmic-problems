package main

func commonFactors(a int, b int) (result int) {
	if a > b {
		a, b = b, a
	}

	for i := 2; i <= a; i++ {
		if a%i == 0 && b%i == 0 {
			result++
		}
	}

	return result + 1
}
