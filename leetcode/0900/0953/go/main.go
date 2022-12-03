package main

func isAlienSorted(words []string, order string) bool {
	letterWeights := getLetterWeights(order)

	for i := 1; i < len(words); i++ {
		if !check(words[i-1], words[i], letterWeights) {
			return false
		}
	}

	return true
}

func check(first, second string, letterWeights []int) bool {
	stop := min(len(first), len(second))
	for i := 0; i < stop; i++ {
		if letterWeights[first[i]] == letterWeights[second[i]] {
			continue
		}

		return letterWeights[first[i]] < letterWeights[second[i]]
	}

	return len(first) <= len(second)
}

func getLetterWeights(order string) []int {
	result := make([]int, 123)
	for i := 0; i < len(order); i++ {
		result[order[i]] = i
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
