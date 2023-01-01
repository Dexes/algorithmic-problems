package main

func wordPattern(pattern string, s string) bool {
	var (
		start, end     = 0, findSpace(s, 1)
		words, letters = [123]string{}, make(map[string]byte)
		pIndex         = 0
	)

	for start < len(s) && pIndex < len(pattern) {
		word, letter := s[start:end], pattern[pIndex]
		if x := words[letter]; x != "" && x != word {
			return false
		}

		if x := letters[word]; x != 0 && x != letter {
			return false
		}

		words[letter], letters[word] = word, letter
		start, end = end+1, findSpace(s, end+2)
		pIndex++
	}

	return pIndex == len(pattern) && start > len(s)
}

func findSpace(s string, index int) int {
	for ; index < len(s) && s[index] != ' '; index++ {
	}

	return index
}
