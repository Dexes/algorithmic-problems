package main

func removeStars(s string) string {
	bytes, wIndex := []byte(s), 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '*' {
			wIndex--
		} else {
			bytes[wIndex] = bytes[i]
			wIndex++
		}
	}

	return string(bytes[:wIndex])
}
