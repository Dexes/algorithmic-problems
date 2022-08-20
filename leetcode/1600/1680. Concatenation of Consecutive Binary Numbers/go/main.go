package main

const modulo = 1e9 + 7

func concatenatedBinary(n int) int {
	result, length, nextLonger := 1, 2, 4
	for x := 2; x <= n; x++ {
		if x == nextLonger {
			nextLonger <<= 1
			length++
		}

		result = ((result % modulo) << length) | x
	}

	return result % modulo
}
