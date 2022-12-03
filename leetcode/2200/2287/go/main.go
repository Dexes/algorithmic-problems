package main

func rearrangeCharacters(s string, target string) int {
	sLetters, tLetters, min := toLetters(s), toLetters(target), -1
	for i := 0; i < len(sLetters); i++ {
		if tLetters[i] > 0 && (min == -1 || sLetters[i]/tLetters[i] < min) {
			min = sLetters[i] / tLetters[i]
		}
	}

	if min == -1 {
		return 0
	}

	return min
}

func toLetters(s string) []int {
	letters := make([]int, 123)
	for i := 0; i < len(s); i++ {
		letters[s[i]]++
	}

	return letters['a':]
}
