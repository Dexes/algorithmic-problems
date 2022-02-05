package main

var isPrime = []bool{
	false, false, true, true, false, true, false, true, false, false,
	false, true, false, true, false, false, false, true, false, true,
}

func countPrimeSetBits(left int, right int) int {
	result := 0
	for ; left <= right; left++ {
		if isPrime[countBits(left)] {
			result++
		}
	}

	return result
}

func countBits(n int) int {
	result := 0
	for n > 0 {
		result += n & 1
		n >>= 1
	}

	return result
}
