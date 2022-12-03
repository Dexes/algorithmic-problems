package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	indices map[int]int
	nums    []int
	random  *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indices: make(map[int]int),
		nums:    make([]int, 0, 1000),
		random:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (x *RandomizedSet) Insert(val int) bool {
	if _, ok := x.indices[val]; ok {
		return false
	}

	x.indices[val] = len(x.nums)
	x.nums = append(x.nums, val)
	return true
}

func (x *RandomizedSet) Remove(val int) bool {
	var index int
	var ok bool

	if index, ok = x.indices[val]; !ok {
		return false
	}

	rightIndex := len(x.nums) - 1
	if index != rightIndex {
		x.indices[x.nums[rightIndex]] = index
		x.nums[index] = x.nums[rightIndex]
	}

	x.nums = x.nums[:rightIndex]
	delete(x.indices, val)
	return true
}

func (x *RandomizedSet) GetRandom() int {
	return x.nums[x.random.Intn(len(x.nums))]
}
