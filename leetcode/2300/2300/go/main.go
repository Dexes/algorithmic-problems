package main

import "sort"

func successfulPairs(spells, potions []int, success64 int64) []int {
	sort.Ints(potions)
	success := int(success64)

	for i := 0; i < len(spells); i++ {
		target := (success + spells[i] - 1) / spells[i]
		index := sort.Search(len(potions), func(i int) bool { return potions[i] >= target })
		spells[i] = len(potions) - index
	}

	return spells
}
