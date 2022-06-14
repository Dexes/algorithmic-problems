package main

func longestCommonPrefix(data []string) string {
	for charIndex := 0; charIndex < len(data[0]); charIndex++ {
		char := data[0][charIndex]
		for i := 1; i < len(data); i++ {
			if len(data[i]) <= charIndex || data[i][charIndex] != char {
				return data[0][:charIndex]
			}
		}
	}

	return data[0]
}
