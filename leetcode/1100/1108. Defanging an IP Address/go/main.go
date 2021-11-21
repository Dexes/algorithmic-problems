package main

func defangIPaddr(address string) string {
	result := make([]byte, len(address)+6)
	current := 0

	for i := 0; i < len(address); i++ {
		if address[i] != '.' {
			result[current] = address[i]
			current++
			continue
		}

		result[current] = '['
		result[current+1] = '.'
		result[current+2] = ']'
		current += 3
	}

	return string(result)
}
