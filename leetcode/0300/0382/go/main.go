package main

import (
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	data []int
}

func Constructor(head *ListNode) Solution {
	data := make([]int, 0, 10000)
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}

	return Solution{data: data}
}

func (x *Solution) GetRandom() int {
	return x.data[rand.Intn(len(x.data))]
}
