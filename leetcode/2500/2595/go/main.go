package main

func evenOddBit(n int) []int {
	return []int{count(n), count(n >> 1)}
}

func count(n int) (result int) {
	for ; n > 0; n >>= 2 {
		if n&1 == 1 {
			result++
		}
	}

	return result
}
