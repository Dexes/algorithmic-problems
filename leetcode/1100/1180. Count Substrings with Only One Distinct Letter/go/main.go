package main

func countLetters(s string) (result int) {
	for left, right := 0, 1; left < len(s); left = right {
		for right < len(s) && s[left] == s[right] {
			right++
		}

		length := right - left
		result += length * (length + 1) / 2
	}

	return result
}
