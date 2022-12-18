package main

func similarPairs(words []string) (result int) {
	counter := make(map[int]int)
	for i := 0; i < len(words); i++ {
		counter[toMask(words[i])]++
	}

	for _, x := range counter {
		result += (x - 1) * x / 2
	}

	return result
}

func toMask(s string) (result int) {
	for i := 0; i < len(s); i++ {
		result |= 1 << (s[i] - 'a')
	}

	return result
}
