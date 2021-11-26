package main

func sortSentence(s string) string {
	words := make([][]byte, 62)
	wordsCount := 0

	for i := 0; i < len(s); i++ {
		word := make([]byte, 0, len(s)-i)
		for ; i < len(s) && isLetter(s[i]); i++ {
			word = append(word, s[i])
		}

		index := byte(0)
		for ; i < len(s) && isDigit(s[i]); i++ {
			index = index*10 + (s[i] - '0')
		}

		words[index] = word
		wordsCount++
	}

	result := make([]byte, 0, 200)
	for i := 1; i <= wordsCount; i++ {
		result = append(result, words[i]...)
		result = append(result, ' ')
	}

	return string(result[:len(result)-1])
}

func isLetter(b byte) bool {
	return b >= 'A' && b <= 'z'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
