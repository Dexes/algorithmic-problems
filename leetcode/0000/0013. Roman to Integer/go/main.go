package main

func romanToInt(s string) int {
	result := 0
	for current := 0; current < len(s); current++ {
		if next := current + 1; next < len(s) && isDouble(s[current], s[next]) {
			result += toInt(s[next]) - toInt(s[current])
			current += 1
			continue
		}

		result += toInt(s[current])
	}

	return result
}

func isDouble(a, b byte) bool {
	return a == 'I' && (b == 'V' || b == 'X') ||
		a == 'X' && (b == 'L' || b == 'C') ||
		a == 'C' && (b == 'D' || b == 'M')
}

func toInt(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	default:
		return 1000
	}
}
