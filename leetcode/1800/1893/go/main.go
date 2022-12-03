package main

func isCovered(ranges [][]int, left int, right int) bool {
	rng := uint64(0)
	for i := 0; i < len(ranges); i++ {
		rng = fillRange(rng, max(ranges[i][0], left), min(ranges[i][1], right))
	}

	return rng == fillRange(0, left, right)
}

func fillRange(rng uint64, left, right int) uint64 {
	for ; left <= right; left++ {
		rng |= 1 << left
	}

	return rng
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
