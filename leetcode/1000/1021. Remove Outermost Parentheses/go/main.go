package main

func removeOuterParentheses(s string) string {
	result, counter := make([]byte, 0, len(s)), 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			counter++
			if counter > 1 {
				result = append(result, s[i])
			}

			continue
		}

		counter--
		if counter > 0 {
			result = append(result, s[i])
		}
	}

	return string(result)
}
