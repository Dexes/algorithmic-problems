package main

func reversePrefix(word string, ch byte) string {
	right := position(word, ch)
	if right == -1 {
		return word
	}

	border := right / 2
	result := []byte(word)
	for i := 0; i <= border; i++ {
		result[i], result[right-i] = result[right-i], result[i]
	}

	return string(result)
}

func position(word string, ch byte) int {
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			return i
		}
	}

	return -1
}
