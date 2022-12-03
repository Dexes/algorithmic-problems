package main

func countAndSay(n int) string {
	rBytes, wBytes := make([]byte, 0, 4462), make([]byte, 0, 4462)
	rBytes = append(rBytes, 1)

	for i := 1; i < n; i++ {
		rBytes, wBytes = recount(rBytes, wBytes)
	}

	return toString(rBytes)
}

func recount(rBytes, wBytes []byte) ([]byte, []byte) {
	for i, j := 0, 0; i < len(rBytes); i = j {
		for j = i + 1; j < len(rBytes) && rBytes[i] == rBytes[j]; j++ {
		}

		wBytes = append(wBytes, byte(j-i), rBytes[i])
	}

	return wBytes, rBytes[:0]
}

func toString(bytes []byte) string {
	for i := 0; i < len(bytes); i++ {
		bytes[i] += '0'
	}

	return string(bytes)
}
