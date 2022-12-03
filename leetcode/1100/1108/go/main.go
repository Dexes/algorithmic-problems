package main

func defangIPaddr(address string) string {
	result := make([]byte, 0, len(address)+6)
	for i := 0; i < len(address); i++ {
		if address[i] != '.' {
			result = append(result, address[i])
		} else {
			result = append(result, '[', '.', ']')
		}
	}

	return string(result)
}
