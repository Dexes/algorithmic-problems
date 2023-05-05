package main

var isVowel = [123]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func maxVowels(s string, k int) (result int) {
	for i := 0; i < k; i++ {
		if isVowel[s[i]] {
			result++
		}
	}

	for i, current := k, result; i < len(s) && result < k; i++ {
		if isVowel[s[i-k]] == isVowel[s[i]] {
			continue
		}

		if isVowel[s[i]] {
			if current++; current > result {
				result = current
			}

			continue
		}

		current--
	}

	return result
}
