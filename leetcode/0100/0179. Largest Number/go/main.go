package main

import "sort"

func largestNumber(nums []int) string {
	digits := splitNumbers(nums)
	sort.Slice(digits, func(i, j int) bool {
		return compare(
			append(digits[j], digits[i]...),
			append(digits[i], digits[j]...),
		)
	})

	return toString(digits)
}

func compare(a, b []byte) bool {
	for i, j := len(a)-1, len(b)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if a[i] != b[j] {
			return b[j] < a[i]
		}
	}

	return false
}

func splitNumbers(nums []int) [][]byte {
	result := make([][]byte, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = splitNumber(nums[i])
	}

	return result
}

func splitNumber(x int) []byte {
	if x == 0 {
		return []byte{0}
	}

	result := make([]byte, 0, 20)
	for ; x > 0; x /= 10 {
		result = append(result, byte(x%10))
	}

	return result
}

func toString(digits [][]byte) string {
	if lastIndex := len(digits[0]) - 1; digits[0][lastIndex] == 0 {
		return "0"
	}

	capacity := 0
	for i := 0; i < len(digits); i++ {
		capacity += len(digits[i])
	}

	result := make([]byte, 0, capacity)
	for i := 0; i < len(digits); i++ {
		for j := len(digits[i]) - 1; j >= 0; j-- {
			result = append(result, digits[i][j]+'0')
		}
	}

	return string(result)
}
