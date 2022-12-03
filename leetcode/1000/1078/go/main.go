package main

func findOcurrences(text string, first string, second string) []string {
	result := make([]string, 0, 10)
	firstStart, firstEnd := wordIndices(text, 0)
	secondStart, secondEnd := wordIndices(text, firstEnd+1)
	thirdStart, thirdEnd := wordIndices(text, secondEnd+1)

	for thirdEnd-thirdStart > 0 {
		if check(text, firstStart, firstEnd, first) && check(text, secondStart, secondEnd, second) {
			result = append(result, text[thirdStart:thirdEnd])
		}

		firstStart, firstEnd = secondStart, secondEnd
		secondStart, secondEnd = thirdStart, thirdEnd
		thirdStart, thirdEnd = wordIndices(text, secondEnd+1)
	}

	return result
}

func check(text string, start, end int, word string) bool {
	if end-start != len(word) {
		return false
	}

	for i := start; i < end; i++ {
		if text[i] != word[i-start] {
			return false
		}
	}

	return true
}

func wordIndices(text string, index int) (int, int) {
	i := index
	for ; i < len(text) && text[i] != ' '; i++ {
	}

	return index, i
}
