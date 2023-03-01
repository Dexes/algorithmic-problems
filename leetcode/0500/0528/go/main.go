package main

import (
	"math/rand"
	"sort"
)

type Solution []int

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}

	return w
}

func (s Solution) PickIndex() int {
	x := rand.Intn(s[len(s)-1])
	return sort.Search(len(s), func(i int) bool { return s[i] > x })
}
