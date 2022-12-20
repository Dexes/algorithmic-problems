package main

func canVisitAllRooms(rooms [][]int) bool {
	return visit(rooms, 0) == len(rooms)
}

func visit(graph [][]int, node int) (result int) {
	if graph[node] == nil {
		return 0
	}

	nodes := graph[node]
	graph[node] = nil

	for _, next := range nodes {
		result += visit(graph, next)
	}

	return result + 1
}
