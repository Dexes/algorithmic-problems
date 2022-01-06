package main

func replaceDigits(s string) string {
	result := []byte(s)
	for i := 1; i < len(result); i += 2 {
		result[i] = result[i-1] + result[i] - '0'
	}

	return string(result)
}
