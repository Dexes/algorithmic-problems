package main

func binaryGap(n int) int {
	result, counter := 0, 0
	n, _ = findOne(n)

	for n > 0 {
		n, counter = findOne(n)
		if counter > result {
			result = counter
		}
	}

	return result
}

func findOne(n int) (int, int) {
	counter := 1
	for n&1 == 0 {
		n >>= 1
		counter++
	}

	return n >> 1, counter
}
