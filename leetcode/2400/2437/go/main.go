package main

func countTime(time string) int {
	return countHours(time[0], time[1]) * countMinutes(time[3], time[4])
}

func countHours(a, b byte) int {
	switch {
	case a == '?' && b == '?':
		return 24
	case a == '?':
		if b < '4' {
			return 3
		}

		return 2
	case b == '?':
		if a < '2' {
			return 10
		}

		return 4
	default:
		return 1
	}
}

func countMinutes(a, b byte) int {
	switch {
	case a == '?' && b == '?':
		return 60
	case a == '?':
		return 6
	case b == '?':
		return 10
	default:
		return 1
	}
}
