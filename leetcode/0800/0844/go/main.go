package main

func backspaceCompare(s string, t string) bool {
	first, second := clearString(s), clearString(t)
	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func clearString(s string) []byte {
	result := make([]byte, len(s))
	length := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if length > 0 {
				length--
			}

			continue
		}

		result[length] = s[i]
		length++
	}

	return result[:length]
}
