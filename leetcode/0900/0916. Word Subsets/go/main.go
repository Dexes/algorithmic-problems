package main

func wordSubsets(words1 []string, words2 []string) []string {
	set, totalLetters := toSet(words2)

	rIndex, wIndex := 0, 0
	for ; rIndex < len(words1); rIndex++ {
		if isSubset(words1[rIndex], totalLetters, set) {
			words1[wIndex] = words1[rIndex]
			wIndex++
		}
	}

	return words1[:wIndex]
}

func isSubset(word string, totalLetters int, set [26]byte) bool {
	if len(word) < totalLetters {
		return false
	}

	for i := 0; i < len(word); i++ {
		letter := word[i] - 'a'
		if set[letter] > 0 {
			set[letter]--
			totalLetters--
		}
	}

	return totalLetters == 0
}

func toSet(words []string) ([26]byte, int) {
	set, totalLetters := [26]byte{}, 0
	for _, word := range words {
		wordSet := [26]byte{}

		for i := 0; i < len(word); i++ {
			letter := word[i] - 'a'
			if wordSet[letter]++; set[letter] < wordSet[letter] {
				totalLetters += int(wordSet[letter] - set[letter])
				set[letter] = wordSet[letter]
			}
		}
	}

	return set, totalLetters
}
