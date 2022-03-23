package main

func brokenCalc(startValue int, target int) int {
	result := 0
	for target > startValue {
		if target&1 == 1 {
			target++
		} else {
			target /= 2
		}

		result++
	}

	return result + startValue - target
}
