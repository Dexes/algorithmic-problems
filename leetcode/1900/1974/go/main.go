package main

func minTimeToType(word string) int {
	result, pointer := 0, 0
	for i := 0; i < len(word); i++ {
		letter := int(word[i] - 'a')
		result += min(diff(letter, pointer), diff(pointer, letter)) + 1
		pointer = letter
	}

	return result
}

func diff(a, b int) int {
	a -= b
	if a < 0 {
		a += 26
	}

	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
