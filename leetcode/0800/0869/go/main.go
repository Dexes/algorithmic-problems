package main

type digits [10]int

func reorderedPowerOf2(n int) bool {
	nDigits := toDigits(n)
	min, max := truncate(toMinNumber(nDigits)), toMaxNumber(nDigits)
	for i := min; i <= max; i <<= 1 {
		if toDigits(i) == nDigits {
			return true
		}
	}

	return false
}

func toDigits(x int) (result digits) {
	for ; x > 0; x /= 10 {
		result[x%10]++
	}

	return result
}

func toMinNumber(x digits) (result int) {
	for i := 0; i < 10; i++ {
		for j := x[i]; j > 0; j-- {
			result = result*10 + i
		}
	}

	return result
}

func toMaxNumber(x digits) (result int) {
	for i := 9; i >= 0; i-- {
		for j := x[i]; j > 0; j-- {
			result = result*10 + i
		}
	}

	return result
}

func truncate(x int) int {
	for bit := 0; ; bit++ {
		if x >>= 1; x == 0 {
			return 1 << bit
		}
	}
}
