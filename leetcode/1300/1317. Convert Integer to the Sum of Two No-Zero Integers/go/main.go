package main

func getNoZeroIntegers(n int) []int {
	a, b := 1, n-1
	for containsZero(a) || containsZero(b) {
		a, b = a+1, b-1
	}

	return []int{a, b}
}

func containsZero(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return true
		}

		n /= 10
	}

	return false
}
