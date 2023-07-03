package main

func maxProbability(n int, edges [][]int, probabilities []float64, start, end int) float64 {
	graph := toGraph(n, edges, probabilities)

	var (
		source vertex
		heap   = heapify(graph[start])
	)

	probabilities = make([]float64, n)
	for _, e := range graph[start] {
		probabilities[e.id] = e.probability
	}
	graph[start] = nil

	for len(heap) > 0 {
		source, heap = extract(heap)
		for _, destination := range graph[source.id] {
			probability := source.probability * destination.probability
			if probability < probabilities[destination.id] {
				continue
			}

			probabilities[destination.id], destination.probability = probability, probability
			heap = insert(heap, destination)
		}

		graph[source.id] = nil
	}

	return probabilities[end]
}

type vertex struct {
	id          int
	probability float64
}

func toGraph(n int, edges [][]int, probabilities []float64) [][]vertex {
	result := make([][]vertex, n)
	for i, e := range edges {
		if probabilities[i] == 0 {
			continue
		}

		result[e[0]] = append(result[e[0]], vertex{probability: probabilities[i], id: e[1]})
		result[e[1]] = append(result[e[1]], vertex{probability: probabilities[i], id: e[0]})
	}

	return result
}

func heapify(data []vertex) []vertex {
	for i := len(data)/2 - 1; i >= 0; i-- {
		siftDown(data, i)
	}

	return data
}

func extract(heap []vertex) (vertex, []vertex) {
	result := heap[0]
	lastIndex := len(heap) - 1

	heap[0], heap = heap[lastIndex], heap[:lastIndex]
	siftDown(heap, 0)

	return result, heap
}

func insert(heap []vertex, x vertex) []vertex {
	heap = append(heap, x)
	siftUp(heap, len(heap)-1)

	return heap
}

func siftDown(heap []vertex, i int) {
	for cIndex := getChildIndex(heap, i); cIndex > 0 && heap[i].probability < heap[cIndex].probability; cIndex = getChildIndex(heap, i) {
		heap[i], heap[cIndex] = heap[cIndex], heap[i]
		i = cIndex
	}
}

func siftUp(heap []vertex, i int) {
	for pIndex := (i - 1) / 2; heap[pIndex].probability < heap[i].probability; pIndex = (i - 1) / 2 {
		heap[i], heap[pIndex] = heap[pIndex], heap[i]
		i = pIndex
	}
}

func getChildIndex(heap []vertex, i int) int {
	switch lIndex, rIndex := i*2+1, i*2+2; {
	case lIndex >= len(heap):
		return -1
	case rIndex >= len(heap):
		return lIndex
	case heap[rIndex].probability < heap[lIndex].probability:
		return lIndex
	default:
		return rIndex
	}
}
