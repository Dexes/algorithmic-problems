package main

func countValidWords(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		for ; i < len(s) && isSpace(s[i]); i++ {
		}

		isValid, hyphenCount := i < len(s), 0
		for ; i < len(s) && !isSpace(s[i]); i++ {
			if isDigit(s[i]) {
				isValid = false
				break
			}

			if isPunctuation(s[i]) {
				if i+1 < len(s) && !isSpace(s[i+1]) {
					isValid = false
					break
				}
			}

			if isHyphen(s[i]) {
				if hyphenCount == 1 {
					isValid = false
					break
				}

				if i == 0 || i+1 == len(s) {
					isValid = false
					break
				}

				if !isLetter(s[i-1]) || !isLetter(s[i+1]) {
					isValid = false
					break
				}

				hyphenCount++
			}
		}

		if isValid {
			result++
		}

		for ; i < len(s) && !isSpace(s[i]); i++ {
		}
	}

	return result
}

func isSpace(b byte) bool {
	return b == ' '
}

func isHyphen(b byte) bool {
	return b == '-'
}

func isPunctuation(b byte) bool {
	return b == ',' || b == '.' || b == '!'
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
