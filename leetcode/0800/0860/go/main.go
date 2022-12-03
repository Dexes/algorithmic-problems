package main

func lemonadeChange(bills []int) bool {
	fives, dozens := 0, 0
	for i := 0; i < len(bills) && fives >= 0 && dozens >= 0; i++ {
		fives, dozens = collect(bills[i], fives, dozens)
	}

	return fives >= 0 && dozens >= 0
}

func collect(num, fives, dozens int) (int, int) {
	if num == 5 {
		return fives + 1, dozens
	}

	if num == 10 {
		return fives - 1, dozens + 1
	}

	if dozens > 0 {
		return fives - 1, dozens - 1
	}

	return fives - 3, dozens
}
