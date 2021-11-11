package main

func reverse(x int) int {
	result, sign := 0, 1
	if x < 0 {
		x, sign = -x, -sign
	}

	for x > 0 {
		result, x = result*10+x%10, x/10
	}

	result *= sign
	if result < -1<<31 || result > 1<<31-1 {
		return 0
	}

	return result
}
