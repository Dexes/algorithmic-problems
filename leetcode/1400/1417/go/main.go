package main

func reformat(s string) string {
	first, second := split(s)
	if len(second) > len(first) {
		first, second = second, first
	}

	if len(first)-len(second) > 1 {
		return ""
	}

	result := make([]byte, 0, len(s))
	for i := 0; i < len(second); i++ {
		result = append(result, first[i], second[i])
	}

	if len(first) > len(second) {
		result = append(result, first[len(first)-1])
	}

	return string(result)
}

func split(s string) ([]byte, []byte) {
	capacity := len(s)/2 + 1
	letters, digits := make([]byte, 0, capacity), make([]byte, 0, capacity)
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' {
			digits = append(digits, s[i])
		} else {
			letters = append(letters, s[i])
		}
	}

	return letters, digits
}
