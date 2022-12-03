package main

func average(salary []int) float64 {
	min, max, sum := salary[0], salary[0], salary[0]
	for i := 1; i < len(salary); i++ {
		sum += salary[i]

		switch {
		case salary[i] > max:
			max = salary[i]
		case salary[i] < min:
			min = salary[i]
		}
	}

	return float64(sum-min-max) / float64(len(salary)-2)
}
