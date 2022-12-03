package main

func sumOfNumberAndReverse(num int) bool {
	if num == 0 {
		return true
	}

	for x := (num + 1) >> 1; x < num; x++ {
		if x+reverse(x) == num {
			return true
		}
	}

	return false
}

func reverse(x int) (result int) {
	for ; x > 0; x /= 10 {
		result = (result * 10) + (x % 10)
	}

	return result
}
