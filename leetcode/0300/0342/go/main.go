package main

func isPowerOfFour(n int) bool {
	if n <= 0 || (n-1)&n != 0 {
		return false
	}

	zerosNumber := 0
	for n != 1 {
		n >>= 1
		zerosNumber++
	}

	return zerosNumber&1 == 0
}
