package main

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	result, index := 0, 1
	for n > 0 {
		if n&1 == 0 {
			result |= index
		}

		index <<= 1
		n >>= 1
	}

	return result
}
