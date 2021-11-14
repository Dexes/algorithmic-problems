package main

func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	data := []byte(s)

	for left < right {
		for ; left < right; left++ {
			if isVowel(data[left]) {
				break
			}
		}

		for ; right > left; right-- {
			if isVowel(data[right]) {
				break
			}
		}

		if left != right {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}

	return string(data)
}

func isVowel(char uint8) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
		char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
}
