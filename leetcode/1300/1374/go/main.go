package main

func generateTheString(n int) string {
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = 'a'
	}

	if n&1 == 0 {
		result[n-1] = 'b'
	}

	return string(result)
}
