package main

func maximumValue(strs []string) (result int) {
	for i := 0; i < len(strs); i++ {
		if x := value(strs[i]); x > result {
			result = x
		}
	}

	return result
}

func value(s string) (result int) {
	for i := 0; i < len(s); i++ {
		if s[i] > '9' {
			return len(s)
		}

		result = result*10 + int(s[i]-'0')
	}

	return result
}
