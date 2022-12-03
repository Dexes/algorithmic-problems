package main

func minimumMoves(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			result++
			i += 2
		}
	}

	return result
}
