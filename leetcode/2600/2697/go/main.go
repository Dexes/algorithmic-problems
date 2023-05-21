package main

func makeSmallestPalindrome(s string) string {
	bytes := []byte(s)
	for left, right := 0, len(bytes)-1; left < right; left, right = left+1, right-1 {
		if bytes[left] < bytes[right] {
			bytes[right] = bytes[left]
		} else {
			bytes[left] = bytes[right]
		}
	}

	return string(bytes)
}
