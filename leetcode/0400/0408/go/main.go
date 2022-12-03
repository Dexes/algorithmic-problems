package main

func validWordAbbreviation(word string, abbr string) bool {
	if len(word) < len(abbr) {
		return false
	}

	num, wordIndex, abbrIndex := 0, 0, 0
	for abbrIndex < len(abbr) && wordIndex < len(word) {
		if '1' <= abbr[abbrIndex] && abbr[abbrIndex] <= '9' {
			num, abbrIndex = readInt(abbr, abbrIndex)
			wordIndex += num
			continue
		}

		if word[wordIndex] != abbr[abbrIndex] {
			return false
		}

		abbrIndex, wordIndex = abbrIndex+1, wordIndex+1
	}

	return abbrIndex == len(abbr) && wordIndex == len(word)
}

func readInt(s string, i int) (result, length int) {
	for ; i < len(s) && s[i] <= '9'; i++ {
		result = result*10 + int(s[i]-'0')
	}

	return result, i
}
