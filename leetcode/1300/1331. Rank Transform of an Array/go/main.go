package main

import "sort"

type Number struct {
	Index, Value int
}

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	numbers := make([]Number, len(arr))
	for i := 0; i < len(arr); i++ {
		numbers[i] = Number{Index: i, Value: arr[i]}
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].Value < numbers[j].Value
	})

	arr[numbers[0].Index] = 1
	for i, rank := 1, 1; i < len(numbers); i++ {
		if numbers[i].Value != numbers[i-1].Value {
			rank++
		}

		arr[numbers[i].Index] = rank
	}

	return arr
}
