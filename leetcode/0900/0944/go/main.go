package main

func minDeletionSize(strings []string) int {
	result := 0
	for i := 0; i < len(strings[0]); i++ {
		if !check(strings, i) {
			result++
		}
	}

	return result
}

func check(strings []string, col int) bool {
	for i := 1; i < len(strings); i++ {
		if strings[i-1][col] > strings[i][col] {
			return false
		}
	}

	return true
}
