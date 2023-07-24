package main

func splitWordsBySeparator(words []string, separator byte) []string {
	result := make([]string, 0, len(words))
	for _, word := range words {
		result = split(result, word, separator)
	}

	return result
}

func split(result []string, word string, separator byte) []string {
	for i, j := 0, 0; ; j++ {
		if j == len(word) {
			return insert(result, word, i, j)
		}

		if word[j] == separator {
			result = insert(result, word, i, j)
			i = j + 1
		}
	}
}

func insert(result []string, word string, start, end int) []string {
	if start == end {
		return result
	}

	return append(result, word[start:end])
}
