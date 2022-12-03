package main

func getPalindrome(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}

func longestPalindrome(s string) string {
	left, right, maxLeft, maxRight := 0, 0, 0, 0
	for i := 1; i < len(s); i++ {
		left, right = getPalindrome(s, i, i)
		if right-left > maxRight-maxLeft {
			maxLeft, maxRight = left, right
		}

		left, right = getPalindrome(s, i-1, i)
		if right-left > maxRight-maxLeft {
			maxLeft, maxRight = left, right
		}
	}

	return s[maxLeft : maxRight+1]
}
