package main

func sumZero(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}

	n--
	result[n] = result[n] * n / -2

	return result
}
