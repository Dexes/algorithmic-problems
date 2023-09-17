package main

import "sort"

func maximumCoins(heroes []int, monsters []int, coins []int) []int64 {
	indices := make([]int, len(heroes))
	for i := 1; i < len(indices); i++ {
		indices[i] = i
	}

	sort.Sort(Pairs{heroes, indices})
	sort.Sort(Pairs{monsters, coins})

	result := make([]int64, len(heroes))
	hIndex, mIndex, sum := 0, 0, 0
	for ; hIndex < len(heroes); hIndex++ {
		for ; mIndex < len(monsters) && monsters[mIndex] <= heroes[hIndex]; mIndex++ {
			sum += coins[mIndex]
		}

		result[indices[hIndex]] = int64(sum)
	}

	return result
}

type Pairs struct {
	first, second []int
}

func (p Pairs) Len() int {
	return len(p.first)
}

func (p Pairs) Swap(i, j int) {
	p.first[i], p.first[j] = p.first[j], p.first[i]
	p.second[i], p.second[j] = p.second[j], p.second[i]
}

func (p Pairs) Less(i, j int) bool {
	return p.first[i] < p.first[j]
}
