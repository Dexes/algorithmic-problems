package main

func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0, 300)
	lineStart, lineEnd, lineLength := 0, 0, 0
	for i := 0; i < len(words); i = lineEnd {
		lineStart, lineEnd, lineLength = getLineBounds(words, i, maxWidth)
		if lineEnd < len(words) {
			result = append(result, makeLine(words, lineStart, lineEnd, lineLength, maxWidth))
		} else {
			result = append(result, makeLastLine(words, lineStart, maxWidth))
		}
	}

	return result
}

func makeLine(words []string, lineStart, lineEnd, lineLength, maxWidth int) string {
	result, wordsCount := make([]byte, 0, maxWidth), lineEnd-lineStart
	result = append(result, words[lineStart]...)

	if wordsCount == 1 {
		return string(appendSpaces(result, maxWidth-len(result)))
	}

	spaceCount, spaceBuffer := (maxWidth-lineLength)/(wordsCount-1), (maxWidth-lineLength)%(wordsCount-1)
	for i := lineStart + 1; i < lineEnd; i++ {
		result = appendSpaces(result, spaceCount)
		if spaceBuffer > 0 {
			result = append(result, ' ')
			spaceBuffer--
		}

		result = append(result, words[i]...)
	}

	return string(result)
}

func makeLastLine(words []string, lineStart, maxWidth int) string {
	result := make([]byte, 0, maxWidth)
	result = append(result, words[lineStart]...)
	for i := lineStart + 1; i < len(words); i++ {
		result = append(result, ' ')
		result = append(result, words[i]...)
	}

	return string(appendSpaces(result, maxWidth-len(result)))
}

func appendSpaces(line []byte, count int) []byte {
	for i := 0; i < count; i++ {
		line = append(line, ' ')
	}

	return line
}

func getLineBounds(words []string, start, maxWidth int) (int, int, int) {
	length := 0
	for i := start; i < len(words); i++ {
		if length+len(words[i]) > maxWidth {
			return start, i, length - (i - start)
		}

		length += len(words[i]) + 1
	}

	return start, len(words), length - (len(words) - start)
}
