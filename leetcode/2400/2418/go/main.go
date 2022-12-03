package main

func sortPeople(names []string, heights []int) []string {
	quickSort(names, heights, 0, len(heights)-1)
	return names
}

func quickSort(names []string, heights []int, low, high int) {
	if low < high {
		p := partition(names, heights, low, high)
		quickSort(names, heights, low, p-1)
		quickSort(names, heights, p+1, high)
	}
}

func partition(names []string, heights []int, low, high int) int {
	for j := low; j < high; j++ {
		if heights[j] > heights[high] {
			heights[low], heights[j] = heights[j], heights[low]
			names[low], names[j] = names[j], names[low]
			low++
		}
	}

	heights[low], heights[high] = heights[high], heights[low]
	names[low], names[high] = names[high], names[low]
	return low
}
