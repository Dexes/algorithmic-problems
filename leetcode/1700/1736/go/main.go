package main

func maximumTime(time string) string {
	result := []byte(time)
	result[0], result[1] = setHours(result[0], result[1])
	result[3], result[4] = setMinutes(result[3], result[4])

	return string(result)
}

func setHours(a, b byte) (byte, byte) {
	switch {
	case a == '?' && b == '?' || a == '2' && b == '?':
		return '2', '3'
	case a == '?' && b < '4':
		return '2', b
	case a == '?':
		return '1', b
	case b == '?':
		return a, '9'
	default:
		return a, b
	}
}

func setMinutes(a, b byte) (byte, byte) {
	switch {
	case a == '?' && b == '?':
		return '5', '9'
	case a == '?':
		return '5', b
	case b == '?':
		return a, '9'
	default:
		return a, b
	}
}
