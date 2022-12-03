package main

func dietPlanPerformance(calories []int, k int, lower int, upper int) (result int) {
	sum, i := 0, 0
	for ; i < k; i++ {
		sum += calories[i]
	}

	if sum < lower {
		result--
	} else if sum > upper {
		result++
	}

	for ; i < len(calories); i++ {
		sum += calories[i] - calories[i-k]
		if sum < lower {
			result--
		} else if sum > upper {
			result++
		}
	}

	return result
}
