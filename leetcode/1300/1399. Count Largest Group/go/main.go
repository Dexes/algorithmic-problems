package main

func countLargestGroup(n int) int {
	sums := make([]int, 37)
	for i := 1; i <= n; i++ {
		sums[sum(i)]++
	}

	max, count := 0, 1
	for i := 0; i < len(sums); i++ {
		switch {
		case sums[i] > max:
			max, count = sums[i], 1
		case sums[i] == max:
			count++
		}
	}

	return count
}

func sum(n int) int {
	result := 0
	for n > 0 {
		result += n % 10
		n /= 10
	}

	return result
}
