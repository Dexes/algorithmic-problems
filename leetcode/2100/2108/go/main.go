package main

func firstPalindrome(words []string) string {
	for i := 0; i < len(words); i++ {
		if isPalindrome(words[i]) {
			return words[i]
		}
	}

	return ""
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left, right = left+1, right-1
	}

	return true
}
