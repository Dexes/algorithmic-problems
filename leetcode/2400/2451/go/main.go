package main

func oddString(words []string) string {
	for letterIndex := 1; ; letterIndex++ {
		first, second := diff(words[0], letterIndex), diff(words[1], letterIndex)
		if first != second {
			if diff(words[2], letterIndex) == first {
				return words[1]
			}

			return words[0]
		}

		for wordIndex := 2; wordIndex < len(words); wordIndex++ {
			if diff(words[wordIndex], letterIndex) != first {
				return words[wordIndex]
			}
		}
	}
}

func diff(word string, index int) int {
	return int(word[index]) - int(word[index-1])
}
