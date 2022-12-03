package main

func canBeTypedWords(text string, brokenLetters string) int {
	result, broken := 0, toArray(brokenLetters)
	for i := 0; i < len(text); i++ {
		canType := true
		for ; i < len(text) && text[i] != ' '; i++ {
			canType = canType && !broken[text[i]-'a']
		}

		if canType {
			result++
		}
	}

	return result
}

func toArray(s string) []bool {
	result := make([]bool, 26)
	for i := 0; i < len(s); i++ {
		result[s[i]-'a'] = true
	}

	return result
}
