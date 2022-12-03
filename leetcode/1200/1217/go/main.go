package main

func minCostToMoveChips(position []int) int {
	even := 0
	for i := 0; i < len(position); i++ {
		if position[i]&1 == 0 {
			even++
		}
	}

	return min(len(position)-even, even)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
