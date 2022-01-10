package main

func decrypt(code []int, k int) []int {
	if k > 0 {
		return next(code, k)
	}

	if k < 0 {
		return prev(code, -k)
	}

	return zeroes(code)
}

func next(code []int, k int) []int {
	sum, result := 0, make([]int, len(code))
	for i := 0; i < k; i++ {
		sum += code[i]
	}

	for i := 0; i < len(code); i++ {
		sum -= code[i]
		sum += code[(i+k)%len(code)]
		result[i] = sum
	}

	return result
}

func prev(code []int, k int) []int {
	sum, result := 0, make([]int, len(code))
	for i := len(code) - 1; i > len(code)-k-1; i-- {
		sum += code[i]
	}

	for i := len(code) - 1; i >= 0; i-- {
		sum -= code[i]
		if i-k < 0 {
			sum += code[len(code)+i-k]
		} else {
			sum += code[i-k]
		}

		result[i] = sum
	}

	return result
}

func zeroes(code []int) []int {
	for i := 0; i < len(code); i++ {
		code[i] = 0
	}

	return code
}
