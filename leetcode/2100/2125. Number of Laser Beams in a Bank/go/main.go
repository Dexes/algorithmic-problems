package main

func numberOfBeams(bank []string) int {
	prevBits, currentBits, result := 0, 0, 0
	for i := 0; i < len(bank); i++ {
		if currentBits = countBits(bank[i]); currentBits != 0 {
			result += prevBits * currentBits
			prevBits = currentBits
		}
	}

	return result
}

func countBits(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			result++
		}
	}

	return result
}
