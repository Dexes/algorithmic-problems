package main

func closetTarget(words []string, target string, startIndex int) int {
	result := len(words)
	for i := 0; i < len(words); i++ {
		if words[i] != target {
			continue
		}

		distance := abs(startIndex - i)
		result = min(result, distance, len(words)-distance)
	}

	if result == len(words) {
		return -1
	}

	return result
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}

	if b < c {
		return b
	}

	return c
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
