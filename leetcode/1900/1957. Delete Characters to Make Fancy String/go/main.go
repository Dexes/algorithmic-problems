package main

func makeFancyString(s string) string {
	result := make([]byte, 0, len(s))
	lastChar, lastCharCount := byte(0), 0
	for i := 0; i < len(s); i++ {
		if s[i] == lastChar && lastCharCount > 1 {
			continue
		}

		result = append(result, s[i])

		if s[i] != lastChar {
			lastChar, lastCharCount = s[i], 1
		} else {
			lastCharCount++
		}
	}

	return string(result)
}
