package main

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	index := 0

	if isUpperLetter(word[0]) {
		index++
	}

	isUpper := isUpperLetter(word[index])
	for index++; index < len(word); index++ {
		if isUpper != isUpperLetter(word[index]) {
			return false
		}
	}

	return true
}

func isUpperLetter(b byte) bool {
	return 'A' <= b && b <= 'Z'
}
