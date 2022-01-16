package main

func divideString(s string, k int, fill byte) []string {
	resultLength, remainder := len(s)/k, len(s)%k
	result := make([]string, 0, resultLength+1)

	for i := 0; i < resultLength; i++ {
		position := i * k
		result = append(result, s[position:position+k])
	}

	if remainder > 0 {
		last := append(make([]byte, 0, k), s[resultLength*k:]...)
		for i := len(last); i < k; i++ {
			last = append(last, fill)
		}

		result = append(result, string(last))
	}

	return result
}
