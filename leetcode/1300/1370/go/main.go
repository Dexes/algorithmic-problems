package main

func sortString(s string) string {
	result, letters := make([]byte, 0, len(s)), countLetters(s)

	for len(result) < len(s) {
		for i := 0; i <= 25; i++ {
			if letters[i] != 0 {
				result = append(result, byte(i)+'a')
				letters[i]--
			}
		}

		for i := 25; i >= 0; i-- {
			if letters[i] != 0 {
				result = append(result, byte(i)+'a')
				letters[i]--
			}
		}
	}

	return string(result)
}

func countLetters(s string) []int {
	result := make([]int, 26)
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}

	return result
}
