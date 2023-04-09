package main

import "math"

func diagonalPrime(nums [][]int) (result int) {
	for i, last := 0, len(nums)-1; i < len(nums); i++ {
		result = maxPrime(result, nums[i][i])
		result = maxPrime(result, nums[i][last-i])
	}

	return result
}

func maxPrime(a, b int) int {
	if a > b || !isPrime(b) {
		return a
	}

	return b
}

func isPrime(x int) bool {
	if x == 1 {
		return false
	}

	for i := int(math.Sqrt(float64(x))); i > 1; i-- {
		if x%i == 0 {
			return false
		}
	}

	return true
}
