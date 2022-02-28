package main

var (
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	modulo = 1000000007
)

func numPrimeArrangements(n int) int {
	count := countPrimes(n)
	return multiply(factorial(count), factorial(n-count))
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = multiply(result, i)
	}

	return result
}

func multiply(x, y int) int {
	return x * y % modulo
}

func countPrimes(n int) int {
	i := 0
	for ; i < len(primes) && primes[i] <= n; i++ {
	}

	return i
}
