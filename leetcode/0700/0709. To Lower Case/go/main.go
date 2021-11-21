package main

func toLowerCase(s string) string {
	result := []byte(s)
	for i := 0; i < len(result); i++ {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] += 32
		}
	}

	return string(result)
}
