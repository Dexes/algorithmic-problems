package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	return min(forward(distance, start, destination), backward(distance, start, destination))
}

func forward(distance []int, start, destination int) int {
	result := 0
	for start != destination {
		result += distance[start]

		if start++; start == len(distance) {
			start = 0
		}
	}

	return result
}

func backward(distance []int, start, destination int) int {
	result := 0
	for start != destination {
		if start--; start == -1 {
			start = len(distance) - 1
		}

		result += distance[start]
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
