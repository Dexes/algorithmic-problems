package main

func reverseOnlyLetters(s string) string {
	data, left, right := []byte(s), 0, len(s)-1
	for left < right {
		for ; left < right && !isLetter(data[left]); left++ {
		}

		for ; right > left && !isLetter(data[right]); right-- {
		}

		data[left], data[right] = data[right], data[left]
		left, right = left+1, right-1
	}

	return string(data)
}

func isLetter(b byte) bool {
	return 'A' <= b && b <= 'Z' || 'a' <= b && b <= 'z'
}
