package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	data map[int][]int
}

func Constructor(nums []int) Solution {
	data := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		list, ok := data[nums[i]]
		if !ok {
			list = make([]int, 0, 5)
		}

		data[nums[i]] = append(list, i)
	}

	rand.Seed(time.Now().UnixNano())
	return Solution{data: data}
}

func (x *Solution) Pick(target int) int {
	list := x.data[target]
	return list[rand.Intn(len(list))]
}
