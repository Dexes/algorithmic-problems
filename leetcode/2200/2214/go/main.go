package main

func minimumHealth(damage []int, armor int) int64 {
	maxDamage, sum := 0, 1
	for i := 0; i < len(damage); i++ {
		maxDamage = max(maxDamage, damage[i])
		sum += damage[i]
	}

	return int64(sum - min(maxDamage, armor))
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
