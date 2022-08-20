package main

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	if len(items2) > len(items1) {
		items1, items2 = items2, items1
	}

	counter := make(map[int]int, len(items1))
	for i := 0; i < len(items1); i++ {
		counter[items1[i][0]] = items1[i][1]
	}

	for i := 0; i < len(items2); i++ {
		counter[items2[i][0]] += items2[i][1]
	}

	index := 0
	for value, weight := range counter {
		items1 = insert(items1, index, value, weight)
		index++
	}

	sort.Slice(items1, func(i, j int) bool {
		return items1[i][0] < items1[j][0]
	})

	return items1
}

func insert(arr [][]int, index, value, weight int) [][]int {
	if len(arr) == index {
		return append(arr, []int{value, weight})
	}

	arr[index][0], arr[index][1] = value, weight
	return arr
}
