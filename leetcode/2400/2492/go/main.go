package main

func minScore(n int, roads [][]int) int {
	graph := toGraph(n, roads)
	return getMinDistance(1, graph[1][0].Distance, graph)
}

func getMinDistance(source, result int, graph [][]City) int {
	var cities []City
	cities, graph[source] = graph[source], nil

	for _, city := range cities {
		result = getMinDistance(city.ID, min(result, city.Distance), graph)
	}

	return result
}

func toGraph(n int, roads [][]int) [][]City {
	result := make([][]City, n+1)
	for _, edge := range roads {
		insert(result, edge[0], edge[1], edge[2])
		insert(result, edge[1], edge[0], edge[2])
	}

	return result
}

func insert(graph [][]City, src, dst, distance int) {
	if graph[src] == nil {
		graph[src] = make([]City, 0, 4)
	}

	graph[src] = append(graph[src], City{ID: dst, Distance: distance})
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

type City struct {
	ID       int
	Distance int
}
