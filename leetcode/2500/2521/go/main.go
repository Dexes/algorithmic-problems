package main

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

func distinctPrimeFactors(nums []int) int {
	localPrimes, processed := make(map[int]struct{}), make([]bool, 1001)
	for i := 0; i < len(nums); i++ {
		factorizeNumber(nums[i], localPrimes, processed)
	}

	return len(localPrimes)
}

func factorizeNumber(x int, localPrimes map[int]struct{}, processed []bool) {
	if processed[x] {
		return
	}

	for i := 0; i < len(primes) && primes[i] < x; i++ {
		if (x % primes[i]) > 0 {
			continue
		}

		localPrimes[primes[i]] = struct{}{}
		for (x % primes[i]) == 0 {
			if x /= primes[i]; processed[x] {
				return
			}

			processed[x] = true
		}
	}

	if x > 1 {
		localPrimes[x] = struct{}{}
	}
}
