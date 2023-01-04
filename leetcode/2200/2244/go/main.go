package main

func minimumRounds(tasks []int) (result int) {
	frequencies := make(map[int]int)
	for i := 0; i < len(tasks); i++ {
		frequencies[tasks[i]]++
	}

	for _, f := range frequencies {
		if f < 2 {
			return -1
		}

		result += (f + 2) / 3
	}

	return result
}
