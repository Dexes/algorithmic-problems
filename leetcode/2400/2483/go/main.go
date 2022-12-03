package main

func bestClosingTime(customers string) int {
	nCount, yCount := 0, 0
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			yCount++
		}
	}

	min, time := yCount, -1
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'N' {
			nCount++
		} else {
			yCount--
		}

		if x := nCount + yCount; x < min {
			min, time = x, i
		}
	}

	return time + 1
}
