package main

func maxNumberOfBalloons(text string) int {
	letters := map[byte]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}
	for i := 0; i < len(text); i++ {
		if text[i] == 'b' || text[i] == 'a' || text[i] == 'l' || text[i] == 'o' || text[i] == 'n' {
			letters[text[i]]++
		}
	}

	letters['l'] /= 2
	letters['o'] /= 2

	return min(letters)
}

func min(letters map[byte]int) int {
	if len(letters) == 0 {
		return 0
	}

	result := 10000
	for _, count := range letters {
		if count < result {
			result = count
		}
	}

	return result
}
