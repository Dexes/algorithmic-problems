package main

func maximumPopulation(logs [][]int) int {
	counter := make([]int, 101)
	for i := 0; i < len(logs); i++ {
		counter[logs[i][0]-1950]++
		counter[logs[i][1]-1950]--
	}

	result, current, max := 0, 0, 0
	for i := 0; i < len(counter); i++ {
		current += counter[i]
		if current > max {
			result, max = i, current
		}
	}

	return result + 1950
}
