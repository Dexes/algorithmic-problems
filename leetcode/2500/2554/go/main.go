package main

func maxCount(banned []int, n int, maxSum int) (result int) {
	exclude := [10001]bool{}
	for i := 0; i < len(banned); i++ {
		exclude[banned[i]] = true
	}

	for i, sum := 1, 0; i <= n; i++ {
		if exclude[i] {
			continue
		}

		if sum += i; sum > maxSum {
			return result
		}

		result++
	}

	return result
}
