package main

func removeComments(source []string) []string {
	var (
		tail           string
		rIndex, wIndex = 0, 0
	)

	for ; rIndex < len(source); rIndex++ {
		for i := 1; i < len(source[rIndex]); i++ {
			if source[rIndex][i] == '/' && source[rIndex][i-1] == '/' {
				source[rIndex] = source[rIndex][:i-1]
				break
			}

			if source[rIndex][i] == '*' && source[rIndex][i-1] == '/' {
				start := []byte(source[rIndex][:i-1])
				rIndex, tail = removeMultiline(source, rIndex, i+2)
				source[rIndex] = string(append(start, tail...))
				i--
			}
		}

		if len(source[rIndex]) > 0 {
			source[wIndex] = source[rIndex]
			wIndex++
		}
	}

	return source[:wIndex]
}

func removeMultiline(source []string, rIndex, i int) (int, string) {
	for ; rIndex < len(source); rIndex++ {
		for ; i < len(source[rIndex]); i++ {
			if source[rIndex][i] == '/' && source[rIndex][i-1] == '*' {
				if next := i + 1; next < len(source[rIndex]) {
					return rIndex, source[rIndex][next:]
				}

				return rIndex, ""
			}
		}

		i = 1
	}

	return rIndex, ""
}
