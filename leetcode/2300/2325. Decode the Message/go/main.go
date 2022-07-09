package main

func decodeMessage(key string, message string) string {
	result, letters := make([]byte, len(message)), decode(key)
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			result[i] = ' '
			continue
		}

		result[i] = letters[message[i]-'a']
	}

	return string(result)
}

func decode(key string) [26]byte {
	result := [26]byte{}
	for i, actualLetter := 0, byte('a'); actualLetter <= 'z'; i++ {
		if key[i] == ' ' {
			continue
		}

		keyLetter := key[i] - 'a'
		if result[keyLetter] > 0 {
			continue
		}

		result[keyLetter] = actualLetter
		actualLetter++
	}

	return result
}
