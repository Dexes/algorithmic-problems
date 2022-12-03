package main

func minOperations(s string) int {
	first, second := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			first++
		} else {
			second++
		}

		first, second = second, first
	}

	return min(first, second)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
