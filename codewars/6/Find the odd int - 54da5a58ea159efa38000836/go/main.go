package kata

func FindOdd(seq []int) (result int) {
	for _, x := range seq {
		result ^= x
	}

	return result
}
