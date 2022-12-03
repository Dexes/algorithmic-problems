package main

func countBits(n int) []int {
	result := make([]int, n+1)

	var left, right int
	for i := 1; i <= n; i = left {
		result[i] = 1

		left, right = i+1, min(i<<1, n+1)
		for ; left < right; left++ {
			result[left] = result[left-i] + 1
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
