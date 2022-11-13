package main

func canSeePersonsCount(heights []int) []int {
	stack, stackIndex, prevIndex, height := make([]int, len(heights)), -1, 0, 0
	for i := len(heights) - 1; i >= 0; i-- {
		for prevIndex = stackIndex; stackIndex >= 0 && stack[stackIndex] < heights[i]; stackIndex-- {
		}

		height, heights[i] = heights[i], prevIndex-stackIndex
		if stackIndex >= 0 {
			heights[i]++
		}

		stackIndex++
		stack[stackIndex] = height
	}

	return heights
}
