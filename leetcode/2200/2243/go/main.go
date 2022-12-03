package main

func digitSum(s string, k int) string {
	digits := toDigits(s)
	for len(digits) > k {
		wIndex := 0
		for rIndex := 0; rIndex < len(digits); rIndex += k {
			wIndex = group(digits, wIndex, rIndex, k)
		}

		digits = digits[:wIndex]
	}

	return toString(digits)
}

func group(digits []int, wIndex, rIndex, count int) int {
	sum := 0
	for stop := min(len(digits), rIndex+count); rIndex < stop; rIndex++ {
		sum += digits[rIndex]
	}

	switch {
	case sum < 10:
		digits[wIndex] = sum
		return wIndex + 1
	case sum < 100:
		digits[wIndex], digits[wIndex+1] = sum/10, sum%10
		return wIndex + 2
	default:
		digits[wIndex], digits[wIndex+1], digits[wIndex+2] = sum/100, sum%100/10, sum%10
		return wIndex + 3
	}
}

func toDigits(s string) []int {
	result := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		result[i] = int(s[i] - '0')
	}

	return result
}

func toString(digits []int) string {
	result := make([]byte, len(digits))
	for i := 0; i < len(digits); i++ {
		result[i] = byte(digits[i]) + '0'
	}

	return string(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
