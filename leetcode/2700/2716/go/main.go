package main

func minimizedStringLength(s string) (result int) {
	var letters [123]bool
	for i := 0; i < len(s); i++ {
		if letters[s[i]] {
			continue
		}

		letters[s[i]] = true
		result++
	}

	return result
}
