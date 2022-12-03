package main

func mostCommonWord(paragraph string, banned []string) string {
	bannedSet, counter := toSet(banned), make(map[string]int)
	maxCount, result := 0, ""
	currentWordStart := 0

	for i := 0; i <= len(paragraph); i++ {
		if i < len(paragraph) && isLetter(paragraph[i]) {
			continue
		}

		if i-currentWordStart == 0 {
			currentWordStart = i + 1
			continue
		}

		word := toLower(paragraph[currentWordStart:i])
		if bannedSet[word] {
			currentWordStart = i + 1
			continue
		}

		counter[word]++
		if counter[word] > maxCount {
			maxCount, result = counter[word], word
		}

		currentWordStart = i + 1
	}

	return result
}

func toLower(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			result[i] = s[i] + 32
		} else {
			result[i] = s[i]
		}
	}

	return string(result)
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z'
}

func toSet(a []string) map[string]bool {
	result := make(map[string]bool)
	for i := 0; i < len(a); i++ {
		result[a[i]] = true
	}

	return result
}
