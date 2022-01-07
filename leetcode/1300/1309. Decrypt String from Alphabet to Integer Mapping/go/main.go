package main

func freqAlphabets(s string) string {
	var shift byte
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+2] == '#' {
			shift = (s[i]-'0')*10 + (s[i+1] - '0') - 1
			i += 2
		} else {
			shift = s[i] - '0' - 1
		}

		result = append(result, 'a'+shift)
	}

	return string(result)
}
