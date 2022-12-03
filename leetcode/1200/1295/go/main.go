package main

func findNumbers(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if countDigits(nums[i])&1 == 0 {
			result++
		}
	}

	return result
}

func countDigits(n int) int {
	if n < 10 {
		return 1
	}

	if n < 100 {
		return 2
	}

	if n < 1000 {
		return 3
	}

	if n < 10000 {
		return 4
	}

	if n < 100000 {
		return 5
	}

	return 6
}
