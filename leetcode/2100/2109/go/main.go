package main

func addSpaces(s string, spaces []int) string {
	result := make([]byte, 0, len(s)+len(spaces))
	strIndex, spaceIndex := 0, 0

	for ; strIndex < len(s) && spaceIndex < len(spaces); strIndex++ {
		if strIndex == spaces[spaceIndex] {
			result = append(result, ' ')
			spaceIndex++
		}

		result = append(result, s[strIndex])
	}

	return string(append(result, s[strIndex:]...))
}
