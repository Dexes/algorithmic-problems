package main

func removePalindromeSub(s string) int {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return 2
		}

		left, right = left+1, right-1
	}

	return 1
}
