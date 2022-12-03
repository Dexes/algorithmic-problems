package main

func validPalindrome(s string) bool {
	return isPalindrome(s, 0, len(s)-1, 1)
}

func isPalindrome(s string, left, right, errors int) bool {
	for left < right {
		if s[left] == s[right] {
			left, right = left+1, right-1
			continue
		}

		if errors > 0 {
			return isPalindrome(s, left+1, right, errors-1) || isPalindrome(s, left, right-1, errors-1)
		}

		return false
	}

	return true
}
