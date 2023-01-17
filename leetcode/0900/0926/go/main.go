package main

func minFlipsMonoIncr(s string) (result int) {
	for i, ones := 0, 0; i < len(s); i++ {
		if s[i] == '0' {
			result = min(result+1, ones)
		} else {
			ones++
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
