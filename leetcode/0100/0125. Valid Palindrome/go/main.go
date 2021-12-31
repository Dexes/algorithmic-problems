package main

func isPalindrome(s string) bool {
	var leftLetter, rightLetter byte
	left, right := -1, len(s)

	for left < right {
		leftLetter, left = findLetter(s, left+1, 1)
		rightLetter, right = findLetter(s, right-1, -1)
		if leftLetter != rightLetter {
			return false
		}
	}

	return true
}

func findLetter(s string, index int, step int) (byte, int) {
	for {
		if index < 0 || index >= len(s) {
			return 0, index
		}

		if 'A' <= s[index] && s[index] <= 'Z' {
			return s[index] + 32, index
		}

		if 'a' <= s[index] && s[index] <= 'z' || '0' <= s[index] && s[index] <= '9' {
			return s[index], index
		}

		index += step
	}
}
