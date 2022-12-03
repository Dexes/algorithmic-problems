package main

func dailyTemperatures(temperatures []int) []int {
	stack, stackIndex, counter := make([][]int, len(temperatures)), -1, 0
	for i := len(temperatures) - 1; i >= 0; i-- {
		for counter = 1; stackIndex >= 0 && stack[stackIndex][0] <= temperatures[i]; {
			counter += stack[stackIndex][1]
			stackIndex--
		}

		if stackIndex < 0 {
			counter = 0
		}

		stackIndex++
		stack[stackIndex] = []int{temperatures[i], counter}
		temperatures[i] = counter
	}

	return temperatures
}
