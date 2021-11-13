package main

func lengthOfLastWord(s string) int {
	result := 0
	prevIsSpace := true

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			prevIsSpace = true
			continue
		}

		if prevIsSpace {
			result = 1
			prevIsSpace = false
			continue
		}

		result++
	}

	return result
}
