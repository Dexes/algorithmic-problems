package main

import (
	"sort"
)

type SparseVector []Pair

type Pair struct {
	Index int
	Num   int
}

func Constructor(nums []int) SparseVector {
	result := make(SparseVector, 0, len(nums)/2)
	for i, n := range nums {
		if n > 0 {
			result = append(result, Pair{Index: i, Num: n})
		}
	}

	return result
}

func (x SparseVector) dotProduct(y SparseVector) (result int) {
	for i, j := 0, 0; i < len(x) && j < len(y); {
		switch {
		case x[i].Index < y[j].Index:
			i += sort.Search(len(x)-i, func(index int) bool { return x[index+i].Index >= y[j].Index })

		case x[i].Index > y[j].Index:
			j += sort.Search(len(y)-j, func(index int) bool { return y[index+j].Index >= x[i].Index })

		default:
			result += x[i].Num * y[j].Num
			i, j = i+1, j+1
		}
	}

	return result
}
