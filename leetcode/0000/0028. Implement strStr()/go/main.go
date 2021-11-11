package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		success := true
		for j := 0; j < len(needle); j++ {
			if i+j >= len(haystack) {
				success = false
				break
			}

			if haystack[i+j] != needle[j] {
				success = false
				break
			}
		}

		if success {
			return i
		}
	}

	return -1
}
