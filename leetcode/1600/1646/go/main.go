package main

func getMaximumGenerated(n int) int {
	result, data := 0, make([]int, n+2)
	data[0], data[1] = 0, 1

	for i := 0; i < n; i += 2 {
		data[i] = data[i>>1]
		data[i+1] = data[i>>1] + data[i>>1+1]
		result = max(result, data[i], data[i+1])
	}

	return result
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	}

	if b > c {
		return b
	}

	return c
}
