package main

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	index, isUpper := 1, word[0] < 'a'
	if isUpper {
		index, isUpper = 2, word[1] < 'a'
	}

	for ; index < len(word); index++ {
		if isUpper != (word[index] < 'a') {
			return false
		}
	}

	return true
}
