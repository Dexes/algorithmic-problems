package main

func xorOperation(n int, start int) int {
	result := start
	for i := 1; i < n; i++ {
		result ^= start + i<<1
	}

	return result
}
