package main

func nextGreatestLetter(letters []byte, target byte) byte {
	result := byte(255)
	for i := 0; i < len(letters); i++ {
		if letters[i] <= target {
			letters[i] += 'z'
		}

		if letters[i] < result {
			result = letters[i]
		}
	}

	if result > 'z' {
		return result - 'z'
	}

	return result
}
