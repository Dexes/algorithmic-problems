package main

func shiftingLetters(s string, shifts []int) string {
	calcShifts(shifts)

	result := []byte(s)
	for i := 0; i < len(result); i++ {
		result[i] += byte(shifts[i])
		if result[i] > 'z' {
			result[i] -= 26
		}
	}

	return string(result)
}

func calcShifts(shifts []int) {
	shifts[len(shifts)-1] %= 26
	for i := len(shifts) - 2; i >= 0; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
	}
}
