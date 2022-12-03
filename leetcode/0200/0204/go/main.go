package main

func countPrimes(n int) (result int) {
	set := make([]bool, n)
	for i := 2; i < n; i++ {
		if set[i] {
			continue
		}

		for j := i * i; j < n; j += i {
			set[j] = true
		}

		result++
	}

	return result
}
