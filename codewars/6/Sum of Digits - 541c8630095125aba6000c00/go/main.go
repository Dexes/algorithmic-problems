package kata

func DigitalRoot(n int) int {
	for ; n > 9; n = sum(n) {
	}

	return n
}

func sum(n int) (result int) {
	for ; n > 0; n /= 10 {
		result += n % 10
	}

	return result
}
