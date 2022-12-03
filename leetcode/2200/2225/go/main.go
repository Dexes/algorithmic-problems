package main

import "sort"

func findWinners(matches [][]int) [][]int {
	winners, losers := make(map[int]struct{}), make(map[int]int)
	for i := 0; i < len(matches); i++ {
		winners[matches[i][0]] = struct{}{}
		losers[matches[i][1]]++
	}

	result := [][]int{make([]int, 0, len(winners)), make([]int, 0, len(losers)>>1)}
	for id := range winners {
		if losers[id] == 0 {
			result[0] = append(result[0], id)
		}
	}

	for id, loses := range losers {
		if loses == 1 {
			result[1] = append(result[1], id)
		}
	}

	sort.Ints(result[0])
	sort.Ints(result[1])

	return result
}
