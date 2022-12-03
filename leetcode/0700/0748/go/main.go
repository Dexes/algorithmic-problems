package main

func shortestCompletingWord(licensePlate string, words []string) string {
	license, result := parseLicense(licensePlate), ""
	for i := 0; i < len(words); i++ {
		if check(license, words[i]) && (len(result) == 0 || len(words[i]) < len(result)) {
			result = words[i]
		}
	}

	return result
}

func check(license []byte, word string) bool {
	for i := 0; i < len(license); i += 2 {
		letter, count := license[i], license[i+1]
		for j := 0; j < len(word) && count > 0; j++ {
			if word[j] == letter {
				count--
			}
		}

		if count > 0 {
			return false
		}
	}

	return true
}

func parseLicense(licensePlate string) []byte {
	result := countLetters(licensePlate)

	index := 0
	for i := byte('a'); i <= 'z'; i++ {
		if result[i] > 0 {
			result[index] = i
			result[index+1] = result[i]
			index += 2
		}
	}

	return result[:index]
}

func countLetters(s string) []byte {
	result := make([]byte, 123)
	for i := 0; i < len(s); i++ {
		if 'a' <= s[i] && s[i] <= 'z' {
			result[s[i]]++
			continue
		}

		if 'A' <= s[i] && s[i] <= 'Z' {
			result[s[i]+32]++
		}
	}

	return result
}
