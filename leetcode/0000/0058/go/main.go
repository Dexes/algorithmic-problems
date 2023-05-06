package main

func lengthOfLastWord(s string) int {
	left, right := 0, len(s)-1
	for ; s[right] == ' '; right-- {
	}

	for left = right; left >= 0 && s[left] != ' '; left-- {
	}

	return right - left
}
