package main

func findComplement(num int) int {
	result, index := 0, 1
	for num > 0 {
		if num&1 == 0 {
			result |= index
		}

		index <<= 1
		num >>= 1
	}

	return result
}
