package main

func shortestWordDistance(words []string, first, second string) (result int) {
	if first == second {
		return singleWordDistance(words, first)
	}

	fIndex, sIndex := wordIndex(words, first, 0), wordIndex(words, second, 0)
	if fIndex > sIndex {
		first, second = second, first
		fIndex, sIndex = sIndex, fIndex
	}

	if result = sIndex - fIndex; result == 1 {
		return result
	}

	for {
		nextIndex := wordIndex(words, first, fIndex+1)

		if nextIndex == len(words) {
			return result
		}

		if nextIndex < sIndex {
			fIndex = nextIndex
		} else {
			fIndex, sIndex = sIndex, nextIndex
			first, second = second, first
		}

		switch distance := sIndex - fIndex; {
		case distance == 1:
			return 1
		case distance < result:
			result = distance
		}
	}
}

func singleWordDistance(words []string, word string) int {
	prev := wordIndex(words, word, 0)
	current := wordIndex(words, word, prev+1)
	result := len(words)

	for current < len(words) {
		switch distance := current - prev; {
		case distance == 1:
			return 1
		case distance < result:
			result = distance
		}

		prev, current = current, wordIndex(words, word, current+1)
	}

	return result
}

func wordIndex(words []string, word string, index int) int {
	for ; index < len(words) && words[index] != word; index++ {
	}

	return index
}
