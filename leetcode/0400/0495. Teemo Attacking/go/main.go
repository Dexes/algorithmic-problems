package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	result := 0
	for i := 0; i < len(timeSeries); {
		start, end := timeSeries[i], timeSeries[i]+duration
		for i++; i < len(timeSeries) && timeSeries[i] <= end; i++ {
			end = timeSeries[i] + duration
		}

		result += end - start
	}

	return result
}
