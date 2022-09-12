package main

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	toPrefix(nums)
	for i := 0; i < len(queries); i++ {
		queries[i] = sort.Search(len(nums), func(j int) bool {
			return nums[j] > queries[i]
		})
	}

	return queries
}

func toPrefix(nums []int) {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
}
