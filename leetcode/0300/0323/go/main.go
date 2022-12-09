package main

func countComponents(n int, edges [][]int) int {
	groups := make([]int, n)
	for i := 0; i < len(groups); i++ {
		groups[i] = i
	}

	for i := 0; i < len(edges); i++ {
		a, b := sort(merge(groups, edges[i][0]), merge(groups, edges[i][1]))
		if a == b {
			continue
		}

		groups[a], n = b, n-1
	}

	return n
}

func merge(groups []int, i int) int {
	for ; i != groups[i]; i = groups[i] {
		groups[i] = groups[groups[i]]
	}

	return i
}

func sort(a, b int) (int, int) {
	if a < b {
		return a, b
	}

	return b, a
}
