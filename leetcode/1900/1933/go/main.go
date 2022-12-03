package main

func isDecomposable(s string) bool {
	doubleExists, length := false, 0
	for i := 0; i < len(s); i += length {
		length = count(s, i)
		switch x := length % 3; {
		case x == 2:
			if doubleExists {
				return false
			}

			doubleExists = true
		case x > 0:
			return false
		}
	}

	return doubleExists
}

func count(s string, startIndex int) int {
	i := startIndex
	for ; i < len(s) && s[i] == s[startIndex]; i++ {
	}

	return i - startIndex
}
