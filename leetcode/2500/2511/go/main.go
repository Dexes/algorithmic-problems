package main

func captureForts(forts []int) (result int) {
	prev := first(forts)
	for current := prev + 1; current < len(forts); current++ {
		if forts[current] == 0 {
			continue
		}

		if forts[current] != forts[prev] {
			result = max(result, current-prev-1)
		}

		prev = current
	}

	return result
}

func first(forts []int) int {
	for i := 0; i < len(forts); i++ {
		if forts[i] != 0 {
			return i
		}
	}

	return len(forts)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
