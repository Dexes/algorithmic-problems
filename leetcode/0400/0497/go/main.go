package main

import (
	"math/rand"
	"sort"
)

type Solution [][]int

func Constructor(rects [][]int) Solution {
	// rect = [left, bottom, width, points (prefix sum)]

	rects[0][2] -= rects[0][0] - 1
	rects[0][3] = (rects[0][3] - rects[0][1] + 1) * rects[0][2]

	for i := 1; i < len(rects); i++ {
		rects[i][2] -= rects[i][0] - 1
		rects[i][3] = rects[i-1][3] + ((rects[i][3] - rects[i][1] + 1) * rects[i][2])
	}

	return rects
}

func (s Solution) Pick() []int {
	point := rand.Intn(s[len(s)-1][3])
	rect := sort.Search(len(s), func(i int) bool { return s[i][3] > point })
	if rect > 0 {
		point -= s[rect-1][3]
	}

	return []int{s[rect][0] + point%s[rect][2], s[rect][1] + point/s[rect][2]}
}
