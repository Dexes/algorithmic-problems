package kata

func CountSheeps(numbers []bool) (result int) {
	for _, x := range numbers {
		if x {
			result++
		}
	}

	return result
}
