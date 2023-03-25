package main

func countPairs(n int, edges [][]int) int64 {
	graph := toGraph(n, edges)
	result, counter := 0, 0
	for i := 0; i < n; i++ {
		if graph[i] == nil {
			continue
		}

		x := count(graph, i)
		result += counter * x
		counter += x
	}

	return int64(result)
}

func count(graph [][]int, n int) (result int) {
	if graph[n] == nil {
		return 0
	}

	destinations := graph[n]
	graph[n] = nil
	for _, dst := range destinations {
		result += count(graph, dst)
	}

	return result + 1
}

func toGraph(n int, edges [][]int) [][]int {
	result := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		insert(result, edges[i][0], edges[i][1])
		insert(result, edges[i][1], edges[i][0])
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
