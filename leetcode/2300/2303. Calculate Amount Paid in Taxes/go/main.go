package main

func calculateTax(brackets [][]int, incomeInt int) float64 {
	upper, percent, income := float64(brackets[0][0]), float64(brackets[0][1]), float64(incomeInt)
	if income <= upper {
		return income * percent / 100
	}

	result := upper * percent / 100
	income -= upper

	for i := 1; ; i++ {
		upper, percent = float64(brackets[i][0]-brackets[i-1][0]), float64(brackets[i][1])
		if income <= upper {
			return result + income*percent/100
		}

		result += upper * percent / 100
		income -= upper
	}
}
