package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	result, data := make([]string, len(score)), convert(score)
	result[data[0].Index] = "Gold Medal"

	if len(result) > 1 {
		result[data[1].Index] = "Silver Medal"
	}

	if len(result) > 2 {
		result[data[2].Index] = "Bronze Medal"
	}

	for i := 3; i < len(result); i++ {
		result[data[i].Index] = strconv.Itoa(i + 1)
	}

	return result
}

type ScoreInfo struct {
	Score int
	Index int
}

func convert(score []int) []ScoreInfo {
	result := make([]ScoreInfo, len(score))
	for i := 0; i < len(score); i++ {
		result[i] = ScoreInfo{Score: score[i], Index: i}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})

	return result
}
