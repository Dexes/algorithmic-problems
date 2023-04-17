package main

const (
	min = -1 << 31
	max = 1<<31 - 1
)

func reverse(x int) int {
	result, sign := 0, 1
	if x < 0 {
		x, sign = -x, -sign
	}

	for ; x > 0; x /= 10 {
		result = (result * 10) + (x % 10)
	}

	result *= sign
	if result < min || result > max {
		return 0
	}

	return result
}
