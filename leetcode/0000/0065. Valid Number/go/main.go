package main

func isNumber(s string) bool {
	i, isFractional, isValid := readNumber(s, 0)

	if !isValid {
		return false
	}

	if i == len(s) {
		return true
	}

	i, isFractional, isValid = readNumber(s, i+1)
	return isValid && !isFractional && i == len(s)
}

func readNumber(s string, i int) (endIndex int, isFractional bool, isValid bool) {
	if i == len(s) {
		return 0, false, false
	}

	if s[i] == '-' || s[i] == '+' {
		i++
	}

	startIndex, dotExists := i, false
	for ; i < len(s) && s[i] != 'e' && s[i] != 'E'; i++ {
		if s[i] == '.' {
			if dotExists {
				return 0, false, false
			}

			dotExists = true
			continue
		}

		if s[i] < '0' || s[i] > '9' {
			return 0, false, false
		}
	}

	if dotExists {
		startIndex++
	}

	return i, dotExists, i > startIndex
}
