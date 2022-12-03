package main

import (
	"sort"
)

func findOriginalArray(changed []int) []int {
	if len(changed)&1 == 1 {
		return []int{}
	}

	sort.Ints(changed)

	m := make(map[int]int)
	wIndex := len(changed) - 1
	for i := len(changed) - 1; i >= 0; i-- {
		double := changed[i] << 1
		count, exists := m[double]

		if !exists {
			m[changed[i]]++
			continue
		}

		changed[wIndex] = changed[i]
		wIndex--

		if count == 1 {
			delete(m, double)
		} else {
			m[double] = count - 1
		}
	}

	if wIndex != len(changed)/2-1 {
		return []int{}
	}

	return changed[wIndex+1:]
}
