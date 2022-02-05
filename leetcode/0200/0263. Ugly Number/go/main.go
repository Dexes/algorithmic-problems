package main

func isUgly(n int) bool {
	for n > 1 {
		switch 0 {
		case n % 2:
			n /= 2
		case n % 3:
			n /= 3
		case n % 5:
			n /= 5
		default:
			return false
		}
	}

	return n == 1
}
