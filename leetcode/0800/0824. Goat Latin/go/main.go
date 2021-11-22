package main

func toGoatLatin(sentence string) string {
	result := make([]byte, 0, 3200)
	letter, lettersInEnd := byte(0), 2

	for i := 0; i < len(sentence); i++ {
		result = append(result, ' ')

		if !isVowel(sentence[i]) {
			i, letter = i+1, sentence[i]
		}

		for ; i < len(sentence) && sentence[i] != ' '; i++ {
			result = append(result, sentence[i])
		}

		if letter != 0 {
			result = append(result, letter)
			letter = 0
		}

		result = append(result, 'm')
		for j := 0; j < lettersInEnd; j++ {
			result = append(result, 'a')
		}
		lettersInEnd++
	}

	return string(result[1:])
}

func isVowel(char uint8) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
		char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
}
