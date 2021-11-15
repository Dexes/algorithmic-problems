package main

func addStrings(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	result := []byte(a)
	digit, buffer := byte(0), byte(0)
	firstIndex, secondIndex := len(a)-1, len(b)-1

	for firstIndex >= 0 && (secondIndex >= 0 || buffer > 0) {
		if secondIndex < 0 {
			digit = 0
		} else {
			digit = b[secondIndex] - '0'
		}

		result[firstIndex] += digit + buffer
		if result[firstIndex] > '9' {
			result[firstIndex] -= 10
			buffer = 1
		} else {
			buffer = 0
		}

		firstIndex--
		secondIndex--
	}

	if buffer == 0 {
		return string(result)
	}

	return string(append([]byte{'1'}, result...))
}

func main() {
	addStrings("9", "1")
}
