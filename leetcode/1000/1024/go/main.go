package main

func videoStitching(clips [][]int, time int) (result int) {
	graph := make([]int, time+1)
	for i := 0; i < len(clips); i++ {
		if clips[i][0] > time {
			continue
		}

		graph[clips[i][0]] = max(graph[clips[i][0]], clips[i][1])
	}

	currentTime, reachedTime, maxTime := 0, 0, 0
	for reachedTime < time {
		for ; currentTime <= reachedTime; currentTime++ {
			maxTime = max(maxTime, graph[currentTime])
		}

		if reachedTime == maxTime {
			return -1
		}

		reachedTime = maxTime
		result++
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
