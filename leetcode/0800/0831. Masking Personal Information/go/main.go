package main

func maskPII(s string) string {
	if s[0] > '9' {
		return maskEmail(s)
	}

	return maskPhone(s)
}

func maskEmail(s string) string {
	i := 2
	for ; s[i] != '@'; i++ {
	}

	result := make([]byte, 0, len(s)+5)
	result = append(result, lower(s[0]), '*', '*', '*', '*', '*', lower(s[i-1]), '@')

	for i++; i < len(s); i++ {
		if s[i] == '.' {
			result = append(result, '.')
			continue
		}

		result = append(result, lower(s[i]))
	}

	return string(result)
}

func maskPhone(s string) string {
	digits := make([]byte, 0, 13)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' {
			digits = append(digits, s[i])
		}
	}

	switch len(digits) {
	case 13:
		return string([]byte{'+', '*', '*', '*', '-', '*', '*', '*', '-', '*', '*', '*', '-', digits[9], digits[10], digits[11], digits[12]})
	case 12:
		return string([]byte{'+', '*', '*', '-', '*', '*', '*', '-', '*', '*', '*', '-', digits[8], digits[9], digits[10], digits[11]})
	case 11:
		return string([]byte{'+', '*', '-', '*', '*', '*', '-', '*', '*', '*', '-', digits[7], digits[8], digits[9], digits[10]})
	default:
		return string([]byte{'*', '*', '*', '-', '*', '*', '*', '-', digits[6], digits[7], digits[8], digits[9]})
	}
}

func lower(b byte) byte {
	if b < 'a' {
		return b + 32
	}

	return b
}
