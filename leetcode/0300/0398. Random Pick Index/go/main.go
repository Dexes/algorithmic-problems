package main

import (
	"math/rand"
)

type Solution struct {
	data map[int][]int
}

func Constructor(nums []int) Solution {
	data := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		data[nums[i]] = append(data[nums[i]], i)
	}

	return Solution{data: data}
}

func (x *Solution) Pick(target int) int {
	list := x.data[target]
	return list[rand.Intn(len(list))]
}
