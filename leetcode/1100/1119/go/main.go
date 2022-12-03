package main

func removeVowels(s string) string {
	bytes, rIndex, wIndex := []byte(s), 0, 0
	for ; rIndex < len(bytes); rIndex++ {
		if isVowel(bytes[rIndex]) {
			continue
		}

		bytes[wIndex], wIndex = bytes[rIndex], wIndex+1
	}

	return string(bytes[:wIndex])
}

func isVowel(char uint8) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}
