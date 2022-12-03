package main

func sortSentence(s string) string {
	words := getWords(s)
	result := append(make([]byte, 0, len(s)), words[0]...)

	for i := 1; i < len(words); i++ {
		result = append(result, ' ')
		result = append(result, words[i]...)
	}

	return string(result)
}

func getWords(s string) []string {
	words, wordsNumber := make([]string, 9), byte(0)
	for start, end := 0, findDigit(s, 0); start < len(s); start, end = end+2, findDigit(s, end+3) {
		words[s[end]-'1'] = s[start:end]
		if s[end] > wordsNumber {
			wordsNumber = s[end]
		}
	}

	return words[:wordsNumber-'0']
}

func findDigit(s string, i int) int {
	for ; i < len(s) && s[i] > '9'; i++ {
	}

	return i
}
