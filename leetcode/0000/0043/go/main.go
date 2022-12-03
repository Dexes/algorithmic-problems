package main

func multiply(first string, second string) string {
	result := make([]byte, len(first)+len(second))
	for i := len(first) - 1; i >= 0; i-- {
		for j := len(second) - 1; j >= 0; j-- {
			resultIndex := i + j + 1
			result[resultIndex] += (first[i] - '0') * (second[j] - '0')
			if result[resultIndex] >= 10 {
				result[resultIndex-1] += result[resultIndex] / 10
				result[resultIndex] %= 10
			}
		}
	}

	start := 0
	for ; start < len(result)-1 && result[start] == 0; start++ {
	}

	for i := start; i < len(result); i++ {
		result[i] += '0'
	}

	return string(result[start:])
}
