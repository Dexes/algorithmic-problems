package main

func hardestWorker(_ int, logs [][]int) int {
	result, maxTime := logs[0][0], logs[0][1]
	for i := 1; i < len(logs); i++ {
		time := logs[i][1] - logs[i-1][1]

		if time > maxTime || time == maxTime && logs[i][0] < result {
			result, maxTime = logs[i][0], time
		}
	}

	return result
}
