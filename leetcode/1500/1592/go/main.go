package main

func reorderSpaces(text string) string {
	words, counter := getWordsAndCountSpaces(text)
	spaceCount, spaceBuffer := getSpacesInfo(words, counter)

	result := make([]byte, 0, len(text))
	result = append(result, text[words[0][0]:words[0][1]]...)
	for i := 1; i < len(words); i++ {
		result = appendSpaces(result, spaceCount)
		result = append(result, text[words[i][0]:words[i][1]]...)
	}

	return string(appendSpaces(result, spaceBuffer))
}

func getSpacesInfo(words [][]int, counter int) (int, int) {
	if len(words) == 1 {
		return 0, counter
	}

	return counter / (len(words) - 1), counter % (len(words) - 1)
}

func getWordsAndCountSpaces(text string) ([][]int, int) {
	words, counter := make([][]int, 0, 50), 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			counter++
			continue
		}

		wordStart := i
		for ; i < len(text) && isLetter(text[i]); i++ {
		}

		words, i = append(words, []int{wordStart, i}), i-1
	}

	return words, counter
}

func appendSpaces(line []byte, count int) []byte {
	for i := 0; i < count; i++ {
		line = append(line, ' ')
	}

	return line
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z'
}
