package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	original []int
	shuffled []int
	random   *rand.Rand
}

func Constructor(nums []int) Solution {
	shuffled := make([]int, len(nums))
	copy(shuffled, nums)

	return Solution{original: nums, shuffled: shuffled, random: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (x *Solution) Reset() []int {
	return x.original
}

func (x *Solution) Shuffle() []int {
	result, random := x.shuffled, x.random
	for i := 0; i < len(result); i++ {
		j := random.Intn(len(result))
		result[i], result[j] = result[j], result[i]
	}

	return result
}
