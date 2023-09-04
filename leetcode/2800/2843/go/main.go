package main

func countSymmetricIntegers(low int, high int) (result int) {
	if low < 10 {
		low = 11
	}

	twoDigits := low < 100
	for x := low; x <= high; x, twoDigits = increment(x, twoDigits) {
		if twoDigits {
			if x/10 == x%10 {
				result++
			}
			continue
		}

		if a, b := x/100, x%100; (a/10 + a%10) == (b/10 + b%10) {
			result++
		}
	}

	return result
}

func increment(x int, twoDigits bool) (int, bool) {
	if 99 <= x && x < 1000 {
		return 1000, false
	}

	return x + 1, twoDigits
}
