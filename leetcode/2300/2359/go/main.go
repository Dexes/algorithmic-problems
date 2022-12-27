package main

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	distances1, distances2 := toDistances(edges, node1), toDistances(edges, node2)
	result, distance := -1, len(edges)+1
	for i := 0; i < len(edges); i++ {
		if distances1[i] == 0 || distances2[i] == 0 {
			continue
		}

		if d := max(distances1[i], distances2[i]); d < distance {
			result, distance = i, d
		}
	}

	return result
}

func toDistances(edges []int, node int) []int {
	result, distance := make([]int, len(edges)), 1
	for ; node != -1 && result[node] == 0; node = edges[node] {
		result[node] = distance
		distance++
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
