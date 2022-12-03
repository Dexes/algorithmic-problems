package main

import (
	"sort"
)

type Number struct {
	Value int
	Index int
}

func maxSubsequence(nums []int, k int) []int {
	numbers := getNumbers(nums)
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].Value > numbers[j].Value
	})

	numbers = numbers[:k]
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].Index < numbers[j].Index
	})

	for i := 0; i < k; i++ {
		nums[i] = numbers[i].Value
	}

	return nums[:k]
}

func getNumbers(nums []int) []*Number {
	result := make([]*Number, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = &Number{Value: nums[i], Index: i}
	}

	return result
}
