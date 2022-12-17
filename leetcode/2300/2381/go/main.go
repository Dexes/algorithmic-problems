package main

func shiftingLetters(s string, shiftsList [][]int) string {
	shifts := compressShifts(s, shiftsList)
	result := make([]byte, len(s))
	shift := 0

	for i := 0; i < len(s); i++ {
		shift = (shift + shifts[i]) % 26

		switch result[i] = s[i] + byte(shift); {
		case result[i] < 'a':
			result[i] += 26
		case result[i] > 'z':
			result[i] -= 26
		}
	}

	return string(result)
}

func compressShifts(s string, shifts [][]int) []int {
	result := make([]int, len(s)+1)
	for _, shift := range shifts {
		start, end := shift[0], shift[1]+1
		if shift[2] == 0 {
			result[start], result[end] = result[start]-1, result[end]+1
		} else {
			result[start], result[end] = result[start]+1, result[end]-1
		}
	}

	return result
}
