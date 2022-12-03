package main

func capitalizeTitle(title string) string {
	result := make([]byte, len(title))
	for i := 0; i < len(title); i++ {
		if i+1 >= len(title) || i+2 >= len(title) || title[i+1] == ' ' || title[i+2] == ' ' {
			result[i] = toLower(title[i])
		} else {
			result[i] = toUpper(title[i])
		}

		for i++; i < len(title) && title[i] != ' '; i++ {
			result[i] = toLower(title[i])
		}

		if i < len(title) {
			result[i] = ' '
		}
	}

	return string(result)
}

func toLower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + 32
	}

	return b
}

func toUpper(b byte) byte {
	if 'a' <= b && b <= 'z' {
		return b - 32
	}

	return b
}
