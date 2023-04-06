package main

func findBuildings(heights []int) []int {
	wIndex := len(heights) - 1
	max := heights[wIndex]
	heights[wIndex] = wIndex

	for i := wIndex - 1; i >= 0; i-- {
		if heights[i] > max {
			wIndex--
			heights[wIndex], max = i, heights[i]
		}
	}

	return heights[wIndex:]
}
