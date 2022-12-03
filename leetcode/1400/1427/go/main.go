package main

func stringShift(s string, shifts [][]int) string {
	shift := getShift(shifts) % len(s)
	if shift == 0 {
		return s
	}

	offset := (len(s) - shift) % len(s)
	return s[offset:] + s[:offset]
}

func getShift(shifts [][]int) (result int) {
	for i := 0; i < len(shifts); i++ {
		if shifts[i][0] == 0 {
			result -= shifts[i][1]
		} else {
			result += shifts[i][1]
		}
	}

	return result
}
