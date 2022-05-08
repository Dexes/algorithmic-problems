package amin

func largestGoodInteger(num string) string {
	result := byte(0)

	for i := 2; i < len(num); i++ {
		if num[i] == num[i-2] && num[i] == num[i-1] && num[i] > result {
			result = num[i]
		}
	}

	return toString(result)
}

func toString(b byte) string {
	if b == 0 {
		return ""
	}

	return string([]byte{b, b, b})
}
