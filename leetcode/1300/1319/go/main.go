package main

func makeConnected(n int, connections [][]int) int {
	if len(connections)+1 < n {
		return -1
	}

	graph, clusters := toGraph(n, connections), -1
	for i := 0; i < n; i++ {
		if graph[i] == nil {
			continue
		}

		clusters++
		walk(graph, i)
	}

	return clusters
}

func walk(graph [][]int, n int) {
	if graph[n] == nil {
		return
	}

	destinations := graph[n]
	graph[n] = nil
	for _, dst := range destinations {
		walk(graph, dst)
	}
}

func toGraph(n int, connections [][]int) [][]int {
	result := make([][]int, n)
	for i := 0; i < len(connections); i++ {
		insert(result, connections[i][0], connections[i][1])
		insert(result, connections[i][1], connections[i][0])
	}

	empty := make([]int, 0, 0)
	for i := 0; i < n; i++ {
		if result[i] == nil {
			result[i] = empty
		}
	}

	return result
}

func insert(graph [][]int, a, b int) {
	if graph[a] == nil {
		graph[a] = make([]int, 0, 4)
	}

	graph[a] = append(graph[a], b)
}
