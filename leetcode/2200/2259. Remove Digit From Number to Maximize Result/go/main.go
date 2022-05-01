package main

func removeDigit(number string, digit byte) string {
	rightIndex := 0
	for i := 0; i < len(number); i++ {
		if number[i] != digit {
			continue
		}

		if i+1 < len(number) && number[i+1] > digit {
			return number[:i] + number[i+1:]
		}

		rightIndex = i
	}

	return number[:rightIndex] + number[rightIndex+1:]
}
