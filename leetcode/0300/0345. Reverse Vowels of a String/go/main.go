package main

func reverseVowels(s string) string {
	data := []byte(s)
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		for ; left < right && !isVowel(data[left]); left++ {
		}

		for ; right > left && !isVowel(data[right]); right-- {
		}

		data[left], data[right] = data[right], data[left]
	}

	return string(data)
}

func isVowel(char uint8) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
		char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
}
