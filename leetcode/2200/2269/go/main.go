package main

func divisorSubstrings(num, k int) (result int) {
	n, x, factor := readInt(num, k)
	for {
		if x > 0 && num%x == 0 {
			result++
		}

		if n == 0 {
			return result
		}

		n, x = n/10, (x/10)+(n%10)*factor
	}
}

func readInt(num, k int) (int, int, int) {
	result, factor := 0, 1
	for ; k > 0; k-- {
		result += (num % 10) * factor
		factor *= 10
		num /= 10
	}

	return num, result, factor / 10
}
