package main

func numberOfEmployeesWhoMetTarget(hours []int, target int) (result int) {
	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			result++
		}
	}

	return result
}
