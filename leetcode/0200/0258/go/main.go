package main

func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	if x := num % 9; x > 0 {
		return x
	}

	return 9
}
