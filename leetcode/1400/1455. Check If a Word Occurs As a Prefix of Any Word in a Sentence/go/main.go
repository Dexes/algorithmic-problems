package main

func isPrefixOfWord(sentence string, searchWord string) int {
	wordIndex := 1
	for i := 0; i < len(sentence); i = nextWord(sentence, i) {
		if check(sentence, searchWord, i) {
			return wordIndex
		}

		wordIndex++
	}

	return -1
}

func check(sentence string, searchWord string, index int) bool {
	for i := 0; i < len(searchWord) && index < len(sentence); i, index = i+1, index+1 {
		if searchWord[i] != sentence[index] {
			return false
		}
	}

	return true
}

func nextWord(sentence string, index int) int {
	for ; index < len(sentence) && sentence[index] != ' '; index++ {
	}

	return index + 1
}
