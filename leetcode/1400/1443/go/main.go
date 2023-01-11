package main

func minTime(n int, edges [][]int, hasApple []bool) int {
	return dfs(toGraph(n, edges), 0, -1, hasApple)
}

func dfs(graph [][]int, node, parent int, hasApple []bool) (result int) {
	for _, child := range graph[node] {
		if child == parent {
			continue
		}

		x := dfs(graph, child, node, hasApple)
		result += x
		if x > 0 || hasApple[child] {
			result += 2
		}
	}

	return result
}

func toGraph(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		insert(graph, edge[0], edge[1])
		insert(graph, edge[1], edge[0])
	}

	return graph
}

func insert(graph [][]int, u, v int) {
	if graph[u] == nil {
		graph[u] = make([]int, 0, 8)
	}

	graph[u] = append(graph[u], v)
}
