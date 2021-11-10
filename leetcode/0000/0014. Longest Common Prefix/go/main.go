package main

func longestCommonPrefix(data []string) string {
	dataLength := len(data)
	firstWordLength := len(data[0])

	for charIndex := 0; charIndex < firstWordLength; charIndex++ {
		char := data[0][charIndex]
		for i := 1; i < dataLength; i++ {
			if len(data[i]) <= charIndex || data[i][charIndex] != char {
				return data[0][:charIndex]
			}
		}
	}

	return data[0]
}
