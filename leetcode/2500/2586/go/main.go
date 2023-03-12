package main

func vowelStrings(words []string, left int, right int) (result int) {
	for ; left <= right; left++ {
		if x := len(words[left]) - 1; isVowel(words[left][0]) && isVowel(words[left][x]) {
			result++
		}
	}

	return result
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
