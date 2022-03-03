package main

import (
	"sort"
)

type Row struct {
	Index, Soldiers int
}

func kWeakestRows(mat [][]int, k int) []int {
	rows := toRows(mat)
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Soldiers == rows[j].Soldiers {
			return rows[i].Index < rows[j].Index
		}

		return rows[i].Soldiers < rows[j].Soldiers
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = rows[i].Index
	}

	return result
}

func toRows(a [][]int) []Row {
	result := make([]Row, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = Row{Index: i, Soldiers: countSoldiers(a[i])}
	}

	return result
}

func countSoldiers(a []int) int {
	left, right := 0, len(a)-1
	if a[left] == 0 {
		return 0
	}

	if a[right] == 1 {
		return len(a)
	}

	for {
		middle := left + (right-left)/2
		switch {
		case a[middle] == 1:
			left = middle + 1
		case a[middle-1] == 0:
			right = middle - 1
		default:
			return middle
		}
	}
}
