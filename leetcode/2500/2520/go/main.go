package main

func countDigits(num int) (result int) {
	for x := num; x > 0; x /= 10 {
		if num%(x%10) == 0 {
			result++
		}
	}

	return result
}
