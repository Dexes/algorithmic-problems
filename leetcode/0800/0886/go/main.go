package main

func possibleBipartition(n int, dislikes [][]int) bool {
	graph, teams := toGraph(n, dislikes), make([]int, n+1)
	for x := 1; x <= n; x++ {
		if !mark(x, 1, graph, teams) {
			return false
		}
	}

	return true
}

func mark(x, team int, graph [][]int, teams []int) bool {
	if teams[x] != 0 {
		return true
	}

	teams[x] = team
	for _, y := range graph[x] {
		if teams[y] == team || !mark(y, -team, graph, teams) {
			return false
		}
	}

	return true
}

func toGraph(n int, dislikes [][]int) [][]int {
	result := make([][]int, n+1)
	for i := 0; i < len(dislikes); i++ {
		insert(dislikes[i][0], dislikes[i][1], result)
		insert(dislikes[i][1], dislikes[i][0], result)
	}

	return result
}

func insert(a, b int, graph [][]int) {
	if graph[a] == nil {
		graph[a] = make([]int, 0, 4)
	}

	graph[a] = append(graph[a], b)
}
