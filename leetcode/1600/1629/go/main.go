package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	result, maxDuration := keysPressed[0], releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		switch duration := releaseTimes[i] - releaseTimes[i-1]; {
		case duration > maxDuration:
			result, maxDuration = keysPressed[i], duration
		case duration == maxDuration && keysPressed[i] > result:
			result = keysPressed[i]
		}
	}

	return result
}
