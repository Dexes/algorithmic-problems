package main

func isThree(n int) bool {
	if n == 1 {
		return false
	}

	x := sqrt(n)
	return x*x == n && isPrime(x)
}

func isPrime(x int) bool {
	for i := sqrt(x); i > 1; i-- {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func sqrt(x int) int {
	left, right := 0, 100
	for {
		middle := left + (right-left)/2
		begin, end := middle*middle, (middle+1)*(middle+1)

		if begin <= x && x < end {
			return middle
		}

		if x < begin {
			right = middle
		} else {
			left = middle + 1
		}
	}
}
