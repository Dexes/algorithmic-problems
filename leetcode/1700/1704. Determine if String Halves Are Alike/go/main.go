package main

func halvesAreAlike(s string) bool {
	counter, half := 0, len(s)/2
	for i := 0; i < half; i++ {
		if isVowel(s[i]) {
			counter++
		}
	}

	for i := half; i < len(s) && counter >= 0; i++ {
		if isVowel(s[i]) {
			counter--
		}
	}

	return counter == 0
}

func isVowel(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		b += 32
	}

	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
